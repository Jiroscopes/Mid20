<template>
  <div v-if="showModal" class="h-screen w-screen fixed top-0 left-0 z-50 bg-white grid grid-cols-5">
    <div class="md:col-span-3 col-span-full flex">
      <div @click="closeModal" class="hover-text text-darkBlue pl-4 pt-4 top-0 left-0 absolute font-Nunito text-xl flex items-center cursor-pointer"><span class="font-bold text-3xl px-2">X</span> Close</div>
      <div class="flex flex-col justify-center w-full">
        <div class="w-1/2 container-width self-center">
          <h2 class="text-blue font-OpenSans font-bold text-3xl md:text-5xl">
            Contact Us
          </h2>
          <span class="width-change w-48 bg-orange block h-1 mb-8 rounded-full"></span>
          <ModalContactForm />
        </div>
      </div>
    </div>
    <div class="md:col-span-2 col-span-full bg-blue flex items-center">
      <div class="grid grid-cols-6 xl:grid-cols-4">
        <!-- Quick Reply -->
        <div class="col-start-override md:col-span-6 xl:col-span-2 col-span-4 grid grid-cols-5 my-5">
          <img class="col-span-1 self-center" src="/quick-reply.svg" alt="Quick Reply"> 
          <div class="col-span-4 col-start-2">
            <h4 class="text-orange font-Nunito font-bold text-2xl">Quick reply</h4>
            <p class="text-white font-Nunito text-sm">We will get back to you within 24 hours to get to know each other as soon as possible!</p>
          </div>
        </div>
        <!-- Build A Plan -->
        <div class="col-start-override md:col-span-6 xl:col-span-2 col-span-4 grid grid-cols-5 my-5">
          <img class="col-span-1 self-center" src="/build-a-plan.svg" alt="Build A Plan"> 
          <div class="col-span-4 col-start-2">
            <h4 class="text-orange font-Nunito font-bold text-2xl">Build a plan</h4>
            <p class="text-white font-Nunito text-sm">Let’s work together building a plan that synergizes with your goals.</p>
          </div>
        </div>
        <!-- Get Started -->
        <div class="col-start-override md:col-span-6 xl:col-span-2 col-span-4 grid grid-cols-5 my-5">
          <img class="col-span-1 self-center" src="/get-started.svg" alt="Get Started"> 
          <div class="col-span-4 col-start-2">
            <h4 class="text-orange font-Nunito font-bold text-2xl">Get Started</h4>
            <p class="text-white font-Nunito text-sm">Execute a strategic plan with our expertise.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import { EventBus } from '@/event-bus'
import Component from 'vue-class-component'

@Component
export default class ContactModal extends Vue {

  showModal = false;
  callCounter = 0;

  // Lifecycle hook
  created() {
    // Listen for global event
    EventBus.$on('modal-open', () => {
      this.openModal();
    })
  }

  openModal() {
    this.showModal = !this.showModal;
    // Prevent scrolling
    document.querySelector('body').style.overflow = 'hidden';
  }

  closeModal() {
    this.showModal = !this.showModal;
    document.querySelector('body').style.overflow = 'visible';
  }
}
</script>

<style>
.col-start-override {
  grid-column-start: 2;
}

/* For some reason this is required. Tailwind md:w-48 doesnt work right */
@media (max-width: 768px) {
  .width-change {
    @apply w-24;
  }

  .container-width {
    width: 75%;
  }
}
.hover-text {
  transform: scale(1);
  transition: all 0.3s ease;
}
.hover-text:hover {
  transform: scale(0.9);
}
</style>