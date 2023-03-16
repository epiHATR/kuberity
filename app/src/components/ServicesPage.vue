<template>
  <div v-if="isLoading" class="col align-self-center">
    <content-loader
      viewBox="0 0 1024 100"
      :speed="1"
      primaryColor="#f3f3f3"
      secondaryColor="#ecebeb"
    >
    </content-loader>
  </div>

  <div v-if="!isLoading">
    <div class="mb-3 breadcums">
      <h3>K8S services</h3>
      <BreadcumsComp />
    </div>

    <table class="table table-hover">
      <thead>
        <tr>
          <th scope="col">NAMESPACE</th>
          <th scope="col">NAME</th>
          <th scope="col">TYPE</th>
          <th scope="col">CLUSTER-IP</th>
          <th scope="col">EXTERNAL-IP</th>
          <th scope="col">PORT</th>
          <th scope="col">SELCTORS</th>
          <th scope="col">AGE</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in services" :key="item.metadata.uid">
          <td>{{ item.metadata.namespace }}</td>
          <td>{{ item.metadata.name }}</td>
          <td>{{ item.spec.type }}</td>
          <td>{{ item.spec.clusterIP }}</td>
          <td>{{ item.status.loadBalancer.ingress != null ? item.status.loadBalancer.ingress[0].ip : "<none>" }}</td>
          <td>
            {{ stringHelper.portFormat(item.spec.ports) }}
          </td>
          <td>
            <ul>
              <li v-for="(value, key) in item.spec.selector" v-bind:key="key">
                {{ stringHelper.replace(key) }}={{ value }}
              </li>
            </ul>
          </td>
          <td>
            <Timeago
              :datetime="new Date(item.metadata.creationTimestamp)"
              :long="false"
            />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>


<script setup>
import { storeToRefs } from "pinia";
import { Timeago } from "vue2-timeago";
import stringHelper from "@/helpers/stringHelper";
import { ContentLoader } from "vue-content-loader";
import { useAppStore } from "../store/app.store";

import BreadcumsComp from "../components/shared/BreadcumsComp.vue";
const { services,isLoading  } = storeToRefs(useAppStore());
const { fetchServices } = useAppStore();

fetchServices();
</script>

<style>
table thead th {
  font-size: 10px !important;
}

table tbody tr td ul {
  padding: 0px !important;
  margin: 0px !important;
}

table tbody tr td ul li {
  list-style: none !important;
  display: block;
  float: left;
  margin-right: 10px;
}
</style>
