package helpers

import (
	"context"
	"encoding/json"
	"log"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var ClientSet kubernetes.Clientset

func SetSecretData(clientset *kubernetes.Clientset, secret *v1.Secret, dataKey string, dataValue string) *v1.Secret {
	secretClient := clientset.CoreV1().Secrets(secret.Namespace)

	updSec := v1.Secret{
		ObjectMeta: secret.ObjectMeta,
		Data: map[string][]byte{
			dataKey: []byte(dataValue),
		},
	}

	payloadBytes, err := json.Marshal(updSec)
	if err != nil {
		panic(err)
	}

	if _, err = secretClient.Patch(context.TODO(), secret.Name, types.StrategicMergePatchType, payloadBytes, metav1.PatchOptions{}); err != nil {
		panic(err)
	}

	sec, err := secretClient.Get(context.TODO(), secret.Name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return sec
}

func BuildSecretClientSet() *kubernetes.Clientset {
	//use the current context in kubeconfig
	config, err := rest.InClusterConfig()

	if err != nil {
		log.Println("could not find any kubeconfig incluster instance, trying parsing local kubeconfig file.")
		home := homedir.HomeDir()
		kubeconfig := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ClientSet := clientset
	return ClientSet
}

func GetAppSecret(clientSet *kubernetes.Clientset, secretName string, namespace string) *v1.Secret {
	log.Println("getting secrets for initializing app")

	secretClient := clientSet.CoreV1().Secrets(namespace)
	secret, err := secretClient.Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return secret
}

func InitializeSecrets(clientSet *kubernetes.Clientset, secretName string, namespace string) {
	log.Print("initialize secrets for kuberity app")
	secret := GetAppSecret(clientSet, secretName, namespace)

	// verify if admin.password has been provided in secret during installation
	adminPassword := string(secret.Data["admin.password"])
	if len(adminPassword) > 0 {
		log.Printf("running kuberity with admin.password configured in secret %s\n", secretName)
	} else {
		log.Println("no admin password created during installation, creating new random one")
		randomPassword := RandomString(12, "")
		secret = SetSecretData(clientSet, secret, "admin.password", randomPassword)
		log.Printf("running kuberity with default admin.password in secret %s", secretName)
	}

	// verify if admin.jwtTokenKey has been provided in secret during installation
	adminJwtToken := string(secret.Data["jwtTokenKey"])
	if len(adminJwtToken) > 0 {
		log.Printf("running kuberity with jwtTokenKey configured in secret %s\n", secretName)
	} else {
		log.Println("no jwtTokenKey created during installation, creating new random one")
		randomJWT := RandomString(18, "base64")
		secret = SetSecretData(clientSet, secret, "jwtTokenKey", randomJWT)
		log.Printf("running kuberity with default jwtTokenKey in secret %s\n", secretName)
	}
}
