<template>
  <div class="py-20">
    <h1 v-if="error.statusCode === 404" class="text-center block font-OpenSans text-6xl text-blue font-bold">Whoa! Wrong turn.</h1>
    <h1 v-else class="text-center block font-OpenSans text-6xl text-blue font-bold">An error occurred</h1>
    <div class="text-center">
        <NuxtLink class="text-DarkBlue underline font-OpenSans text-lg" to="/">Home page</NuxtLink>
    </div>
  </div>
</template>

<script>
  export default {
    props: ['error'],
    layout: 'error' // you can set a custom layout for the error page
  }
</script>