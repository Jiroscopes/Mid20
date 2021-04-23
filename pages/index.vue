<template>
    <div id="home">
        <Navbar />
        <SectionsHero />
        <SectionsWhatWeDo />
        <SectionsTestimonials />        
        <SectionsTech />
        <img class="absolute right-0 circleLines xl:top-circleLines xl:block hidden" src="/lines_1.png" alt="Circle Pattern">

        <!-- <img class="circle2 testimonial-circles z-10 absolute h-circleMD lg:h-full xl:top-circle2XL lg:top-circle2LG md:top-circle2MD md:block hidden" src="/circle2.png" alt="Circle Pattern"> -->

        <!-- Contact Section -->

        <section id="contact" class="relative grid grid-flow-row grid-rows-1 auto-rows-min md:px-20 px-4 py-24 bg-blue">
            <div class="row-start-1 grid grid-flow-row auto-rows-min grid-cols-8">
                <h2 class="lg:col-span-3 col-span-full text-white font-OpenSans text-3xl font-bold">Want to work together? <br> Let us know how we can help!</h2>
                <span class="row-start-2 md:col-span-3 col-span-full bg-orange block h-1 my-4"></span>
            </div>
            <div class="flex justify-center py-24 z-10">
                <CustomForm />
            </div>
            <p class="absolute bottom-0 left-0 text-white font-Nunito p-3 z-10">	&#169; Mid-Twenty LLC 2020</p>
            <img class="absolute md:block bottom-0 right-0" src="/morecircles.png" alt="Bottom Circles">
        </section>
        <!-- <SectionsLandingContact /> -->
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component';

@Component
export default class index extends Vue {
    data() {
        return {
        }
    }
}
</script>

<style scoped>

.py-60 {
    padding-top: 13rem;
    padding-bottom: 13rem;
}

.circleLines {
    @media (max-width: 1600px) {
        top: 1750px;
    }
}

</style>
