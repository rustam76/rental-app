<template>
  <div class="flex flex-col min-h-screen ">
    <!-- Header -->
    <Header/>

    <!-- Main Content -->
    <main class="flex-grow p-6 container mx-auto">
     <RouterView />
    </main>

    <!-- Footer -->
   <Footer/>
  </div>
</template>

<script setup lang="ts">
import Header from "@/components/layout/header/Header.vue";
import Footer from "@/components/layout/footer/Footer.vue";
</script>
