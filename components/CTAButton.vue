<template>
    <a href="/#contact" v-if="invert" class="hover:bg-orange lg:inline rounded-full py-2 px-8 hover:text-white font-bold font-Nunito bg-transparent text-orange border-orange border-2 transition-all duration-200 ease-in hidden">Get in touch</a>
    <a href="/#contact" v-else class="bg-orange rounded-full py-3 px-10 text-white font-bold font-Nunito mt-8  border-2 transition-all duration-200 ease-in hover:bg-transparent hover:text-orange hover:border-orange">Get in touch</a>
</template>
</div>

<script>
import Vue from 'vue'
import Component from 'vue-class-component'

const CTAButtonProps = Vue.extend({

    props: {
        invert: Boolean,
    }

})

@Component
export default class CTAButton extends CTAButtonProps {

}
</script>

<style scoped>

</style>