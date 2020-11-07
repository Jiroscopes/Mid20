<template>
    <form action="" class="grid w-full md:w-3/5 auto-rows-min">
        <div class="md:row-start-1 grid md:grid-cols-2 auto-rows-min">
            <div class="my-0 md:mx-12 row-start-1">
                <label class="font-Nunito font-bold block text-white" for="first_name">First Name:</label>
                <input class="p-4 h-8 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="first_name">
            </div>
            <div class="my-5 md:my-0 md:mx-12 row-start-2 md:row-start-1">
                <label class="font-Nunito font-bold block text-white" for="business_name">Business Name:</label>
                <input class="p-4 h-8 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="business_name">
            </div>
        </div>
        <div class="row-start-2 md:my-6">
            <div class="md:my-0 md:mx-12">
                <label class="font-Nunito font-bold block text-white" for="email">Email:</label>
                <input class="p-4 h-8 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="email">
            </div>
        </div>
        <div class="row-start-3 my-6">
            <div class="md:mx-12">
                <label class="font-Nunito font-bold block text-white" for="comments">Comments:</label>
                <textarea class="p-4 h-48 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="comments"></textarea>
            </div>
        </div>
        <div class="row-start-4 my-6 flex justify-center">
            <button class="font-Nunito font-bold bg-orange sm:bg-transparent text-white sm:text-orange border-2 border-solid border-orange px-16 py-2 rounded-3xl hover:bg-orange hover:text-white transition duration-500 ease " type="submit">Submit</button>
        </div>
    </form>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

export default class Form extends Vue {


}
</script>

<style>
html {
  scroll-behavior: smooth;
}
</style>
