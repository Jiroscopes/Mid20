<template>
     <p class="text-orange hover:bg-opacity-50 transition duration-200 ease-in px-6 py-2 bg-orange bg-opacity-25 font-OpenSans font-semibold rounded-full text-center"> {{ text }} </p>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

const TechProps = Vue.extend({
     props: {
        text: String,
    }
})

@Component
export default class TechBubble extends TechProps {

}

</script>

<style>

</style>
