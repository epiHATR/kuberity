<template>
    <div>
        <router-link :to="$route.matched[0].path">{{ $route.matched[0].name }}</router-link> > <router-link :to="$route.path">{{$route.name}}</router-link>
    </div>
</template>

<script>
export default {
  name: "BreadcumsComp"
};
</script>