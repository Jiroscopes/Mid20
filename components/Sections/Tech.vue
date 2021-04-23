<template>
  <section class="grid grid-flow-row grid-rows-1 auto-rows-min md:px-32 px-4 py-24 bg-white">
    <div class="row-start-1 grid grid-flow-row auto-rows-min grid-cols-8">
        <h2 class="lg:col-span-full col-span-full text-blue font-OpenSans text-3xl font-bold">Leveraging the best technologies for the job.</h2>
        <span class="row-start-2 col-span-2 bg-orange block h-1 mt-2 mb-4 rounded"></span>
    </div>
    <div class="relative z-20 row-start-2">
        <h3 class="text-orange text-xl font-OpenSans font-semibold">Languages:</h3>
        <div class="grid grid-flow-row xl:grid-cols-7 lg:grid-cols-3 sm:grid-cols-2 grid-cols-1 sm:gap-8 gap-0" >
            <div class="col-span-1 my-4 md:px-4">
                <TechBubble text="JavaScript" />
            </div>
            <div class="bubble col-span-1 my-4 md:px-4">
                <TechBubble text="PHP" />
            </div> 
            <div class="bubble col-span-1 my-4 md:px-4">
                <TechBubble text="C++" />
            </div> 
            <div class="bubble col-span-1 my-4 md:px-4">
                <TechBubble text="Go" />
            </div> 
            <div class="bubble col-span-1 my-4 md:px-4">
                <TechBubble text="HTML/CSS" />
            </div> 
        </div>
    </div>
    <div class="relative z-20 row-start-3">
        <h3 class="text-orange text-xl font-OpenSans font-semibold">Frameworks:</h3>
        <div class="grid grid-flow-row xl:grid-cols-7 lg:grid-cols-3 sm:grid-cols-2 grid-cols-1 sm:gap-8 gap-0" >
            <div class="bubble col-span-1 my-4 md:px-4 min-w-min">
                <TechBubble text="Vue.js" />
            </div>
            <div class="bubble col-span-1 my-4 md:px-4">
                <TechBubble text="React.js" />
            </div> 
            <div class="bubble col-span-1 my-4 md:px-4">
                <TechBubble text="Bootstrap" />
            </div> 
            <div class="bubble col-span-1 my-4 md:px-4">
                <TechBubble text="Express.js" />
            </div>
        </div>
    </div>
  </section>
</template>

<script>
export default {

}
</script>

<style scoped>
.bubble {
    min-width: fit-content;
}
</style>