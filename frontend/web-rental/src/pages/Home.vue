<script setup lang="ts">

import About from '@/pages/About.vue';
import Contact from '@/pages/Contact.vue';
import Explore from '@/pages/Explore.vue';
</script>

<template>
    <section id="hero">
        <div class="w-full ">
            <img src="https://placehold.co/1920x800" alt="" class="rounded-4xl w-full object-cover">
        </div>
    </section>
    <section id="explore" class="mt-3">
        <Explore/>
    </section>
    <section id="about" class="mt-3">
        <About/>
    </section>
    <section id="contact">
        <Contact/>
    </section>
</template>
