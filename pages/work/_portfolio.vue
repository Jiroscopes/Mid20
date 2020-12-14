<template>
    <div id="head">
        <Navbar />
        <div class="bg-white">
            <div class="px-32 py-20 ">
                <img class="shadow-xl" :src="'/' + getPiece() + '/full.png'" :alt="getPiece() + ' Image'">
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