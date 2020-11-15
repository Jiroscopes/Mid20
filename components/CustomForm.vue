<template>
    <form @submit.prevent="checkForm" method="post" class="grid w-full md:w-3/5 auto-rows-min">
        <div class="md:row-start-1 grid md:grid-cols-2 auto-rows-min">
            <div class="my-0 md:mx-12 row-start-1">
                <span v-if="errors[0]" class="text-orange font-bold font-OpenSans">Error: Please enter a valid name</span>
                <label class="font-Nunito font-bold block text-white" for="first_name">* First Name:</label>
                <input v-model="firstName" required class="p-4 h-8 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="first_name">
            </div>
            <div class="my-5 md:my-0 md:mx-12 row-start-2 md:row-start-1">
                <span v-if="errors[1]" class="text-orange font-bold font-OpenSans">Error: Please enter a valid name</span>
                <label class="font-Nunito font-bold block text-white" for="business_name">Business Name:</label>
                <input v-model="businessName" class="p-4 h-8 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="business_name">
            </div>
        </div>
        <div class="row-start-2 md:my-6">
            <div class="md:my-0 md:mx-12">
                <span v-if="errors[2]" class="text-orange font-bold font-OpenSans">Error: Please enter a valid email</span>
                <label class="font-Nunito font-bold block text-white" for="email">* Email:</label>
                <input v-model="email" required class="p-4 h-8 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="email">
            </div>
        </div>
        <div class="row-start-3 my-6">
            <div class="md:mx-12">
                <span v-if="errors[3]" class="text-orange font-bold font-OpenSans">Error: Please enter a comment</span>
                <label class="font-Nunito font-bold block text-white" for="comments">* Comments:</label>
                <textarea v-model="comments" required class="p-4 h-48 rounded-md w-full font-Nunito block text-white bg-darkBlue" type="text" name="comments"></textarea>
            </div>
        </div>
        <div class="row-start-4 my-6 flex justify-center">
            <div>
                <p class="text-center text-white font-Nunito py-4">* Required</p>
                <button class="block font-Nunito font-bold bg-orange sm:bg-transparent text-white sm:text-orange border-2 border-solid border-orange px-16 py-2 rounded-3xl hover:bg-orange hover:text-white transition duration-500 ease " type="submit">Submit</button>
                <span v-if="success">SUCCESS</span>
            </div>
        </div>
    </form>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class Form extends Vue {

    apiURL: string = '/.netlify/functions/lambda';

    firstName: string = '';
    businessName: string = '';
    email: string = '';
    comments: string = '';
    errors: Array<boolean> = [false, false, false, false];
    success: boolean =  false;

    // error = {
    //     errFirstName: false,
    //     errBusinessName: false,
    //     errEmail: false,
    //     errComments: false
    // }
    // firstname, business name, email, comments

    checkForm(e: Event) {

        let arrayOfVal = [this.firstName, this.businessName, this.email, this.comments]

        // For each value
        arrayOfVal.forEach((val, index) => {

            // If not business name, since it's not required
            if (index != 2) {
                if (val == '') {
                    // Update the data using $set because Vue is picky about arrays https://vuejs.org/v2/guide/reactivity.html#For-Arrays
                    this.$set(this.errors, index, true)
                }
            }

            // If email
            if (index === 2) {
                var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                if (!re.test(val)) {
                    this.$set(this.errors, index, true)
                };
            }
        });

        // If any errors exists return
        this.errors.forEach(val => {
            if (val === true) {
                return;
            }
        });

        // Now we can send 
        fetch(this.apiURL, {
            method: 'POST', // or 'PUT'
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({FirstName: this.firstName, businessName: this.businessName, Email: this.email, Comments: this.comments}),
        }).then(async res => {
            if (res.status == 200) {
                this.success = true;
            }
        })
        
    }

}
</script>

<style>
html {
  scroll-behavior: smooth;
}
</style>
