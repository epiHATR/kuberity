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
      <h3>K8S deployments</h3>
      <BreadcumsComp />
    </div>
    <table class="table table-hover">
      <thead>
        <tr>
          <th scope="col">NAMESPACE</th>
          <th scope="col">NAME</th>
          <th scope="col">READY</th>
          <th scope="col">UP-TO-DATE</th>
          <th scope="col">AVAILABLE</th>
          <th scope="col">SELECTORS</th>
          <th scope="col">IMAGE</th>
          <th scope="col">AGE</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in deployments" :key="item.metadata.uid">
          <td>{{ item.metadata.namespace }}</td>
          <td>{{ item.metadata.name }}</td>
          <td>{{ item.status.readyReplicas }}/{{ item.status.replicas }}</td>
          <td>{{ item.status.updatedReplicas }}</td>
          <td>{{ item.status.availableReplicas }}</td>
          <td>
            <ul>
              <li
                v-for="(value, key) in item.spec.selector.matchLabels"
                v-bind:key="key"
              >
                {{ stringHelper.replace(key) }}={{ value }}
              </li>
            </ul>
          </td>
          <td>{{ item.spec.template.spec.containers[0].image }}</td>
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
const { deployments,isLoading  } = storeToRefs(useAppStore());
const { fetchDeployments } = useAppStore();

fetchDeployments();

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
