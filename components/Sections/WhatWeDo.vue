<template>
  <section id="wwd" class="bg-blue grid grid-flow auto-rows-max py-24 md:px-32 px-4">
    <div class="lg:grid lg:grid-cols-10 gap-0 z-20 relative">
        <div class="lg:col-start-1 lg:col-end-5 lg:col-span-5 col-start-1 col-end-10 col-span-10 pb-12 lg:pb-0">
            <h2 class="text-white font-OpenSans font-bold xl:text-6xl text-5xl">What we do</h2>
            <span class="w-48 bg-orange block h-1 mb-8 rounded-full"></span>
            <p class="text-white font-Nunito w-4/5">
                We create high quality web based applications to give your business everything it needs in todays marketplace. 
                The way we do business online has changed for good, just offering services online is not enough anymore, it matters how. 
                Companies that realize the change and can start generating more value will go farther.
                <br />
                <br />
                  We’ll help you grow your online presence. 
                We are committed to understanding and researching your business so that we can serve your business the best way possible.
            </p>
        </div>
        <div class="grid lg:gap-6 lg:grid-cols-6 lg:col-start-5 col-start-1 col-end-11">
            <div @mouseover="triggerShadow('web-design')" id="web-design" class="overflow-hidden col-span-10 md:col-span-5 xl:col-span-3 py-8 lg:py-0 max-h-80 transition-all duration-200 ease-in rounded-3xl wwd-shadow">
                <ServiceBox image="/web_icon.png" title="Website Design" text="Custom website designs that help you make your mark on the web. We design each website to maximize the effectiveness of your company. We want to help you make your online presence a delight for your customers." />
            </div>
            <div @mouseover="triggerShadow('web-development')" id="web-development" class="overflow-hidden col-span-10 md:col-span-5 xl:col-span-3 py-8 lg:py-0 max-h-80 transition-all duration-200 ease-in rounded-3xl">
                <ServiceBox image="/web_dev.png" title="Web Development" text="Build your dream website with completely custom experiences for your business. We can build your online presense to suit your business’s needs, no templates." />
            </div>
            <div @mouseover="triggerShadow('software-development')" id="software-development" class="overflow-hidden col-span-10 md:col-span-5 xl:col-span-3 py-8 lg:py-0 max-h-80 transition-all duration-200 ease-in rounded-3xl">
                <ServiceBox image="/cube.png" title="Software Development" text="From automation, apps,  data analysis, and software integration we can develop it all. We can create any tools your buisness needs to track metrics or automate tasks to make your business run smoother and grow faster." />
            </div>
        </div>
    </div>
  </section>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component';

@Component
export default class WhatWeDo extends Vue {
  data() {
    return {
      activeBox: 'Website Design',
    }
  }

  triggerShadow(id: string): void {
    let serviceBoxes = document.querySelectorAll('.wwd-shadow');

    serviceBoxes.forEach(el => {
        el.classList.remove('wwd-shadow');
    });

    let elem = document.getElementById(id);
    elem?.classList.add('wwd-shadow');
  }
}
</script>

<style>
.wwd-shadow {
    @media(min-width: 1024px) {
        box-shadow: 6px 6px 17px 3px rgba(7, 13, 89, 0.4);
    }
}
</style>