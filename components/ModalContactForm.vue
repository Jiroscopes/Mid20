<template>
    <div>
        <div v-if="success" class="loader flex items-baseline justify-center"><h2 class="text-blue text-4xl">Thank you!</h2></div>
        <div v-if="showLoader" class="loader flex items-baseline justify-center"><img src="/loader.svg" alt=""></div>
        <form @submit.prevent="checkForm" method="post" class="w-full">
            <div class="">
                <span v-if="errors[0]" class="text-orange font-bold font-OpenSans">Please enter a valid name</span>
                <!-- <label class="font-Nunito font-bold block text-white" for="first_name">* First Name:</label> -->
                <input v-model="firstName" placeholder="Name" required class="p-5 h-8 my-4 rounded w-full font-Nunito block text-darkBlue bg-transparent border-2 border-blue" type="text" name="first_name">
            </div>
            <div class="">
                <span v-if="errors[1]" class="text-orange font-bold font-OpenSans">Please enter a valid business name</span>
                <!-- <label class="font-Nunito font-bold block text-darkBlue" for="first_name">* First Name:</label> -->
                <input v-model="businessName" placeholder="Business Name" class="p-5 h-8 my-4 rounded w-full font-Nunito block text-darkBlue bg-transparent border-2 border-blue" type="text" name="business_name">
            </div>
            <div>
                <span v-if="errors[2]" class="text-orange font-bold font-OpenSans">Please enter a valid email</span>
                <input v-model="email" placeholder="Email" required class="p-5 h-8 my-4 rounded w-full font-Nunito block text-darkBlue bg-transparent border-2 border-blue" type="text" name="email">
            </div>
            <div>
                <span v-if="errors[3]" class="text-orange font-bold font-OpenSans">Please tell us about your project</span>
                <textarea v-model="comments" placeholder="About your project..." required class="p-5 h-48 my-4 rounded w-full font-Nunito block text-darkBlue bg-transparent border-2 border-blue" type="text" name="comments"></textarea>
            </div>
            <div class="my-6 flex justify-center text-center">
                <!-- <span v-if="success" class="text-white text-center text-bold font-OpenSans text-3xl py-8">Message Sent. Thank You!</span> -->
                <!-- <p class="text-center text-white font-Nunito py-4">* Required</p> -->
                <button class="m-0 m-auto block font-Nunito font-bold bg-orange sm:bg-transparent text-white sm:text-orange border-2 border-solid border-orange px-16 py-2 rounded-3xl hover:bg-orange hover:text-white transition duration-500 ease " type="submit">Submit</button>
            </div>
        </form>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

@Component
export default class ModalContactForm extends Vue {
    apiURL: string = '/.netlify/functions/lambda';

    firstName: string = '';
    businessName: string = '';
    email: string = '';
    comments: string = '';
    errors: Array<boolean> = [false, false, false, false];
    success: boolean =  false;
    showLoader = false;

    checkForm(e: Event) {
        this.showLoader =  true;
        let arrayOfVal = [this.firstName, this.businessName, this.email, this.comments];

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
input::placeholder {
    color: #070D59;
}
input::-ms-input-placeholder { /* Internet Explorer 10-11 */
  color: #070D59;
}

input::-ms-input-placeholder { /* Microsoft Edge */
  color: #070D59;
}
textarea::placeholder {
    color: #070D59;
}
textarea::-ms-input-placeholder { /* Internet Explorer 10-11 */
  color: #070D59;
}

textarea::-ms-input-placeholder { /* Microsoft Edge */
  color: #070D59;
}

.loader {
    width: 30%;
    height: 100%;
    min-height: 440px;
    position: absolute;
    background: #F6F5F5;
    z-index: 50;
}

@media (max-width: 768px) {
    .loader {
        width: 75%;
        height: 60%;
    }
}
</style>