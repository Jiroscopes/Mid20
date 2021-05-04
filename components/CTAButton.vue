<template>
    <div @click="emitModalClick" v-if="invert" class="cursor-pointer hover:bg-orange lg:inline-block rounded-full py-2 px-8 hover:text-white font-bold font-Nunito bg-transparent text-orange border-orange border-2 transition-all duration-200 ease-in hidden">Get in touch</div>
    <div @click="emitModalClick" v-else class="cursor-pointer bg-orange rounded-full lg:inline-block py-3 px-10 text-white font-bold font-Nunito mt-8 border-transparent border-2 transition-all duration-200 ease-in hover:bg-transparent hover:text-orange hover:border-orange">Get in touch</div>
</template>
</div>

<script>
import Vue from 'vue'
import Component from 'vue-class-component'
import { EventBus } from '@/event-bus'

const CTAButtonProps = Vue.extend({

    props: {
        invert: Boolean,
    }

})

@Component
export default class CTAButton extends CTAButtonProps {
    emitModalClick() {
        EventBus.$emit('modal-open');
    }
}
</script>

<style scoped>

</style>