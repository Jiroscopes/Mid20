<template>
  <section id="contact" class="relative grid grid-flow-row grid-rows-1 auto-rows-min md:px-32 px-4 py-24 bg-blue">
    <div class="text-center">
        <h3 class="text-white font-OpenSans font-bold">Is software important to your business?</h3>
        <h2 class="text-white font-OpenSans font-bold text-4xl">Build it with Mid Twenty</h2>
        <span class="line bg-orange block h-1 mt-2 mb-4 rounded w-1/4"></span>
    </div>
    <div class="text-center">
        <CTAButton />
    </div>
    <img class="footer-circles z-10 absolute h-circleMD lg:h-full lg:block hidden" src="/footerCircles.png" alt="Circle Pattern">
  </section>
</template>

<script>
export default {

}
</script>

<style scoped>
.line {
  margin: 0 auto;
}
.footer-circles {
    right: 0;
    bottom: 0;
    height: 550px;
}
</style>