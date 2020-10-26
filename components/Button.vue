<template>
    <div id="example-1">
        <button v-on:mouseenter="fuckOff">Add 1</button>
        <p> The button above has been clicked {{ counter }} times.</p>
    </div>
</template>
<script>
export default {
    name: "Button",
    data() {
        return {
            counter: 0
        }
    },
    methods: {
        fuckOff() {
            alert("Fuck off");
        }
    }
}
</script>
<style scoped>

</style>