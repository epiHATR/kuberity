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
      <h3>K8S pods</h3>
      <BreadcumsComp />
    </div>
    <table class="table table-hover">
      <thead>
        <tr>
          <th scope="col">NAMESPACE</th>
          <th scope="col">NAME</th>
          <th scope="col">IMAGE</th>
          <th scope="col">STATUS</th>
          <th scope="col">READY</th>
          <th scope="col">RESTARTS</th>
          <th scope="col">IP</th>
          <th scope="col">NODE</th>
          <th scope="col">AGE</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in pods" :key="item.metadata.uid">
          <td>{{ item.metadata.namespace }}</td>
          <td>{{ item.metadata.name }}</td>
          <td>{{ item.spec.containers[0].image }}</td>
          <td>{{ item.status.phase }}</td>
          <td>{{ statusHelper.getReadyOnTotal(item.status.containerStatuses) }}</td>
          <td>{{ statusHelper.countRestart(item.status.containerStatuses) }}</td>
          <td>{{ item.status.podIP }}</td>
          <td>{{ item.spec.nodeName }}</td>
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
import statusHelper from "@/helpers/statusHelper";
import { storeToRefs } from "pinia";
import { Timeago } from "vue2-timeago";
import { ContentLoader } from "vue-content-loader";
import { useAppStore } from "../store/app.store";

import BreadcumsComp from "../components/shared/BreadcumsComp.vue";
const { pods ,isLoading  } = storeToRefs(useAppStore());
const { fetchPods } = useAppStore();

fetchPods();

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
