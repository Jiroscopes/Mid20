<template>
  <div class="bg-white sm:grid sm:pt-20 pt-10 sm:px-20 px-4">
    <div class="sm:grid grid-rows-auto sm:grid-cols-5 md:pt-48 pt-20">
        <div class="col-span-full landing-text">
            <h1 class="lg:text-6xl sm:text-5xl font-OpenSans text-blue text-3xl font-bold"> Innovative software solutions. </h1>
            <h2 class="lg:text-2xl font-Nunito text-darkBlue font-bold  text-xl"> We build the tools that help you grow.</h2>
            <CTAButton class="inline-block" />
        </div>
        <div class="col-span-3 text-right -mt-32 col-start-3 landing-image">
            <img class="inline" src="/landing-img.png" alt="">
        </div> 
    </div>
    <!-- <div class="flex justify-center pb-8">
        <a href="#wwd"><ScrollIcon /></a>
    </div> -->
  </div>
</template>

<script>
export default {

}
</script>

<style>
.landing-text {
    margin-bottom: -300px;
    z-index: 999;
}

/* @media (max-width: 1880px) { */
    .landing-text {
        margin-bottom: 0;
        text-align: center;
        @apply col-span-full;
    }

    .landing-image {
        @apply col-span-full pt-32;
        text-align: center;
    }
/* } */
</style>