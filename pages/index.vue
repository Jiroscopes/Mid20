<template>
    <div>
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
               <a href="#NextPage"><ScrollIcon /></a>
            </div>
        </div>

        <img class="absolute lg:top-xl md:top-lg md:block hidden" src="/circle1.png" alt="Circle Pattern">


        <!-- Services Section -->


        <section class="bg-blue grid grid-flow auto-rows-max">
            <div class="flex justify-center py-24" >
                <img src="../assets/Icons/WWD.png" alt="What We Do" >
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
            <div class="row-start-2">
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
            <div class="row-start-3">
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
        

        <!-- Recent Work -->


        <section class="bg-white md:px-20 px-4">
            <div class="flex justify-center py-24" >
                <img src="/RW.png" alt="Recent Work" >
            </div>
            <div class="grid grid-cols-4 py-">
                <div class="col-start-2 col-span-2">
                    <LargePortfolio image="/d2c.jpg" title="Testing" type="Website" link="/"/>
                </div>
            </div>
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
