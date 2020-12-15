<template>
    <div id="head" class="bg-white ">
        <Navbar />
        <div class="md:px-32 px-2">
            <div class="py-32 text-center md:text-left">
                <h1 class="text-blue font-bold font-OpenSans md:text-6xl text-3xl">
                    {{capitializeName()}}
                </h1>
                <div class="grid grid-flow-row xl:grid-cols-7 lg:grid-cols-3 sm:grid-cols-2 grid-cols-1 sm:gap-8 gap-0" >
                    <div class="col-span-1 my-4 md:px-2">
                        <TechBubble text="JavaScript" />
                    </div>
                    <div class="col-span-1 my-4 md:px-2">
                        <TechBubble text="PHP" />
                    </div> 
                    <div class="col-span-1 my-4 md:px-2">
                        <TechBubble text="WordPress" />
                    </div>  
                    <div class="col-span-1 my-4 md:px-2">
                        <TechBubble text="HTML/CSS" />
                    </div> 
                </div>
            </div>
            <div class="bg-white">
                <div class=" py-10 ">
                    <img class="shadow-xl" :src="'/' + getPiece() + '/full.png'" :alt="getPiece() + ' Image'">
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component';

@Component
export default class index extends Vue {
    data() {
        return {};
    }

    capitializeName(): string {
        let name = this.getPiece();
        let nameCap = name.charAt(0).toUpperCase() + name.slice(1);
        return nameCap;
    }

    getPiece(): string {
        return this.$route.params.portfolio;
    }
}
</script>

<style>
.py-60 {
    padding-top: 13rem;
    padding-bottom: 13rem;
}
</style>