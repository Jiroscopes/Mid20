<template>
    <nav class="flex bg-white flex-wrap justify-between items-center px-20 py-3 w-full fixed z-30">
        <div class="inline-block">
            <div class="text-base">
                <NuxtLink to="/" class="transition duration-500 font-Nunito font-bold text-xl block lg:inline-block lg:mt-0 text-orange hover:text-darkBlue mr-4">
                    <img class="inline" src="/logoV2.png" alt="Mid Twenty Logo">
                    <span class="text-blue font-bold">Mid Twenty</span>
                </NuxtLink>
            </div>
        </div>
        <!-- Toggle Nav Button -->
        <button @click="toggleNav" class="lg:hidden px-3 py-2 border rounded text-blue border-darkBlue hover:text-darkBlue hover:border-darkBlue">
            <svg class="fill-current h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Menu</title><path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"/></svg>
        </button>
        <!-- Menu Items -->
        <div class=" inline-block flex-grow lg:flex lg:items-center lg:w-auto" v-bind:class="{hidden: isHidden, 'w-full': !isHidden}">
            <div class="ml-auto">
                <NuxtLink to="/#home" v-bind:class="{hidden: isHidden}" class="font-Nunito font-bold text-lg block mt-4 lg:inline-block lg:mt-0 text-blue hover:text-darkBlue mr-4">
                    Home
                </NuxtLink>
                <NuxtLink to="/#wwd" v-bind:class="{hidden: isHidden}" class="font-Nunito font-bold text-lg block mt-4 lg:inline-block lg:mt-0 text-blue hover:text-darkBlue mr-4">
                    Services
                </NuxtLink>
                <!-- <NuxtLink to="/" v-bind:class="{hidden: isHidden}" class="font-Nunito font-bold text-lg block mt-4 lg:inline-block lg:mt-0 text-blue hover:text-darkBlue mr-4">
                    Services
                </NuxtLink> -->
                <NuxtLink to="/#contact" v-bind:class="{hidden: isHidden}" class="font-Nunito font-bold text-lg block mt-4 lg:inline-block lg:mt-0 text-blue hover:text-darkBlue mr-4">
                    Contact Us
                </NuxtLink>
                <!-- <button class="hover:bg-orange lg:inline rounded-full py-2 px-8 hover:text-white font-bold font-Nunito bg-transparent text-orange border-orange border-2 transition-all duration-200 ease-in hidden">
                    Get in touch
                </button> -->
                <CTAButton :invert=true />
            </div>
        </div>
    </nav>
</template>
<style scoped>
</style>
<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class Navbar extends Vue {
   
    isHidden: boolean = true;

    toggleNav(): void {
        this.isHidden = !this.isHidden;
    }


}
</script>
