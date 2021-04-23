<template>
  <section class="bg-darkBlue grid grid-flow-row grid-rows-1 auto-rows-min md:px-32 px-4 py-24">
    <div class="row-start-1 grid grid-flow-row auto-rows-min grid-cols-8">
        <h2 class="lg:col-span-full col-span-full text-white font-OpenSans text-3xl font-bold">What our clients are saying...</h2>
        <span class="row-start-2 col-span-1 bg-orange block h-1 mt-2 mb-4 rounded"></span>
    </div>
    <div class="grid grid-flow-row grid-cols-8 py-10 md:gap-4 gap-6">
      <div class="xl:col-span-2 lg:col-span-3 md:col-span-4 col-span-full">
        <p class="text-white font-Nunito">
          <span class="quote text-orange inline-block">“</span>
          The end result was time saved, money saved, a complete solution, and new knowledge. I would highly recommend Mid Twenty to anyone looking for a talented, patient, and communicative team!
        </p>
        <p class="text-white font-Nunito pt-6">-Thomas Matter</p>
      </div>
      <div class="xl:col-start-4 xl:col-span-2 lg:col-span-3 lg:col-start-5 md:col-span-4 md:col-start-5 col-span-full">
        <p class="text-white font-Nunito">
          <span class="quote text-orange inline-block">“</span>
          Mid-Twenty was responsive, communicative and sharp. A great problem solver! I had hired two freelancers prior to Mid-Twenty and they could not figure out the problem, but Mid-Twenty solved it right away. Would highly recommend!
        </p>
        <p class="text-white font-Nunito pt-6">- Jessica, Truster Building and Fence</p>
      </div>
    </div>
  </section>
</template>

<script>
export default {

}
</script>

<style scoped>
.quote {
  font-size: 8rem;
  vertical-align: top;
  line-height: 78px;
  height: 10px;
  display: inline-block;
}
</style>