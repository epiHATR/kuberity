class stringHelper {
    replace(st) {
        var rep = ""
        var repWith = ""

        if (st.includes("app.kubernetes.io/instance")) {
            rep = "app.kubernetes.io/instance"
            repWith = "instance"
        }
        else if (st.includes("app.kubernetes.io/name")) {
            rep = "app.kubernetes.io/name"
            repWith = "app/name"
        }
        else {
            return st
        }
        const result = st.split(rep).join(repWith)
        return result;
    }

    portFormat(st) {
        const ports = JSON.parse(JSON.stringify(st));
        var output=[]
        ports.forEach(port => {
           output.push(port.port + "/" +port.protocol)
        });

        return output.join(", ")
    }
}

export default new stringHelper();