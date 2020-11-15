<template>
    <div id="home">
        <Navbar />
        <div @mousemove="callRef" class="bg-white grid grid-flow-row auto-rows-max">
            <div class="grid grid-cols-3">
                <div class="lg:col-span-1 lg:pl-10 lg:block col-span-0 hidden pl-0">
                    <Phone ref="phoneRef" />
                </div>
                <div class="lg:col-span-2 lg:text-left lg:py-48 text-center col-span-3 py-24 flex items-center justify-center">
                    <div>
                        <h1 class="lg:text-6xl font-OpenSans text-blue text-4xl font-bold"> Your Digital Agency </h1>
                        <h2 class="lg:text-2xl font-Nunito text-orange text-xl"> Want to customize your web presence and stand out? <br> We can help. </h2>
                    </div>
                </div>
            </div>
            <div class="flex justify-center pb-8">
               <a href="#wwd"><ScrollIcon /></a>
            </div>
        </div>

        <img class="absolute lg:top-xl md:top-lg md:block hidden" src="/circle1.png" alt="Circle Pattern">


        <!-- Services Section -->


        <section id="wwd" class="bg-blue grid grid-flow auto-rows-max">
            <div class="flex justify-center py-24" >
                <img src="/icons/WWD.png" alt="What We Do" >
            </div>
            <div class="grid xl:grid-cols-8 lg:gap-6 xl:gap-0 lg:grid-cols-2 grid-cols-10 sm:px-20 px-4 pb-24">
                <div class="xl:col-start-1 xl:col-end-3 lg:col-span-2 col-span-10 py-8 lg:py-0">
                    <ServiceBox image="/web_icon.png" title="Website Design" text="Custom website designs that help you make your mark on the web. We design each website to maximize the effectiveness of your company. We want to help you make your online presence a delight for your customers." />
                </div>
                <div class="xl:col-start-4 xl:col-end-6 lg:col-span-2 col-span-10 py-8 lg:py-0">
                    <ServiceBox image="/web_dev.png" title="Web Development" text="Build your dream website with completely custom experiences for your business. We can build your online presense to suit your business’s needs, no templates." />
                </div>
                <div class="xl:col-start-7 xl:col-end-9 lg:col-span-2 col-span-10 py-8 lg:py-0">
                    <ServiceBox image="/cube.png" title="Software Development" text="From automation, apps,  data analysis, and software integration we can develop it all. We can create any tools your buisness needs to track metrics or automate tasks to make your business run smoother and grow faster" />
                </div>
            </div>
        </section>


        <!-- Tech Section -->
        

        <section class="grid grid-flow-row grid-rows-1 auto-rows-min md:px-20 px-4 py-24 bg-darkBlue">
            <div class="row-start-1 grid grid-flow-row auto-rows-min grid-cols-8">
                <h2 class="lg:col-span-3 col-span-full text-white font-OpenSans text-3xl font-bold">Leveraging the best <br> technologies for the job.</h2>
                <span class="row-start-2 md:col-span-3 col-span-full bg-orange block h-1 my-4"></span>
            </div>
            <div class="relative z-20 row-start-2">
                <h3 class="text-white text-xl font-OpenSans font-semibold">Languages:</h3>
                <div class="grid grid-flow-row xl:grid-cols-7 lg:grid-cols-3 sm:grid-cols-2 grid-cols-1 sm:gap-8 gap-0" >
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="JavaScript" />
                    </div>
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="PHP" />
                    </div> 
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="C++" />
                    </div> 
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="Go" />
                    </div> 
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="HTML/CSS" />
                    </div> 
                </div>
            </div>
            <div class="relative z-20 row-start-3">
                <h3 class="text-white text-xl font-OpenSans font-semibold">Frameworks:</h3>
                <div class="grid grid-flow-row xl:grid-cols-7 lg:grid-cols-3 sm:grid-cols-2 grid-cols-1 sm:gap-8 gap-0" >
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="Vue.js" />
                    </div>
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="React.js" />
                    </div> 
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="Bootstrap" />
                    </div> 
                    <div class="col-span-1 my-4 md:px-4">
                        <TechBubble text="Express.js" />
                    </div>
                </div>
            </div>
        </section>
        

        <img class="absolute right-0 xl:top-circleLines hidden xl:block hidden" src="/lines_1.png" alt="Circle Pattern">

        <!-- Recent Work -->
        

        <section id="recentWork" class="bg-white md:px-20 py-24 px-4">
            <div class="flex justify-center py-24" >
                <img src="/RW.png" alt="Recent Work" >
            </div>
            <div class="grid grid-cols-6 gap-4 md:grid-rows-2 grid-rows-3">
                <div class="xl:col-start-2 col-span-full xl:col-span-4 row-start-1">
                    <LargePortfolio image="" title="Testing" type="Website" link="/"/>
                </div>
                <div class="row-start-2 xl:col-span-2 xl:col-start-2 col-span-full col-start-1 md:col-span-3">
                    <SmallPortfolio title="Testing Small titles" type="Website" link="/" />
                </div>
                <div class="md:row-start-2 row-start-3 xl:col-span-2 xl:col-start-4 col-span-full col-start-1 md:col-span-3">
                    <SmallPortfolio title="Testing Small" type="Website" link="/" />
                </div>
            </div>
        </section>

        <img class="absolute h-circleMD lg:h-full xl:top-circle2XL lg:top-circle2LG md:top-circle2MD md:block hidden" src="/circle2.png" alt="Circle Pattern">

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
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component';

@Component
export default class index extends Vue {
    data() {
        return {
//            mouseEnter: false,
        }
    }

    callRef(): void {
        //@ts-ignore
        this.$refs.phoneRef.changePhonePos(event.clientX, event.clientY);
    }
}
</script>

<style>
.py-60 {
    padding-top: 13rem;
    padding-bottom: 13rem;
}
</style>
