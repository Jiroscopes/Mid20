<template>
    <div id="home">
        <Navbar />
        <div @mousemove="callRef" class="bg-white grid grid-flow-row auto-rows-max">
            <div class="grid grid-cols-5">
                <div class="lg:col-span-2 lg:pl-10 lg:block col-span-0 hidden pl-0">
                    <Phone ref="phoneRef" />
                </div>
                <div class="lg:col-span-3 lg:text-left lg:pb-48 lg:pt-64 text-left col-span-5 pb-40 flex items-center sm:text-center sm:justify-center">
                    <div>
                        <h1 class="lg:text-6xl font-OpenSans text-blue text-4xl font-bold"> Innovative software solutions. </h1>
                        <h2 class="lg:text-2xl font-Nunito text-darkBlue font-bold  text-xl"> We build the tools that help you grow.</h2>
                        <button class="main-CTA">Get in touch</button>
                    </div>
                </div>
            </div>
            <div class="flex justify-center pb-8">
               <a href="#wwd"><ScrollIcon /></a>
            </div>
        </div>

        <img class="absolute lg:top-xl md:top-lg md:block hidden" src="/circle1.png" alt="Circle Pattern">


        <!-- Services Section -->


        <section id="wwd" class="bg-blue grid grid-flow auto-rows-max py-24">
            <div class="grid xl:grid-cols-10 xl:gap-0 sm:px-20 px-4 pb-24">
                <div class="col-start-2 col-end-5 col-span-5">
                    <h2 class="section-headers">What We Do</h2>
                    <span class="w-48 bg-orange block h-1 my-4 rounded-full"></span>
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
                <div class="grid gap-6 grid-cols-6 col-start-5 col-end-11">
                    <div :class="this.activeBox === 'Website Design' ? 'wwd-shadow' : ''"  class="serviceBox-comp-holder">
                        <ServiceBox class="wwd-shadow" image="/web_icon.png" title="Website Design" text="Custom website designs that help you make your mark on the web. We design each website to maximize the effectiveness of your company. We want to help you make your online presence a delight for your customers." />
                    </div>
                    <div class="serviceBox-comp-holder">
                        <ServiceBox image="/web_dev.png" title="Web Development" text="Build your dream website with completely custom experiences for your business. We can build your online presense to suit your business’s needs, no templates." />
                    </div>
                    <div class="serviceBox-comp-holder">
                        <ServiceBox image="/cube.png" title="Software Development" text="From automation, apps,  data analysis, and software integration we can develop it all. We can create any tools your buisness needs to track metrics or automate tasks to make your business run smoother and grow faster." />
                    </div>
                    <div class="serviceBox-comp-holder">
                        <ServiceBox image="/cube.png" title="Product Development" text="TO DO" />
                    </div>
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
                    <LargePortfolio image="/jiroscopes/screenshot.jpg" title="Jiroscopes.com" type="Website" link="/work/jiroscopes" />
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
            activeBox: "Website Design",
        }
    }
    callRef(): void {
        //@ts-ignore
        this.$refs.phoneRef.changePhonePos(event.clientX, event.clientY);
    }

    // checkMouseEvent(title: string) : void {
    //     console.log(title);
    // }
}
</script>

<style scoped>
.py-60 {
    padding-top: 13rem;
    padding-bottom: 13rem;
}

.main-CTA {
    @apply bg-orange rounded-full py-3 px-10 text-white font-bold font-Nunito mt-8 hover:bg-transparent hover:text-orange hover:border-orange border-2 transition-all duration-200 ease-in ;
}

.serviceBox-comp-holder {
    @apply col-span-3 py-8 lg:py-0 max-h-64 wwd-shadow;
}

.section-headers {
    @apply text-white font-OpenSans font-bold text-6xl;
}

.wwd-shadow:hover {
    box-shadow: 6px 6px 17px 3px rgba(7, 13, 89, 0.4);
}
</style>
