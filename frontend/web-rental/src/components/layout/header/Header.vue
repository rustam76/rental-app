<template>
  <header class="p-4">
    <nav class="flex container mx-auto items-center justify-between">
      <!-- <div v-on:click="open = !open"  class="lg:hidden">
            <AlignLeft  />
        </div> -->
      <div class="flex items-center gap-2">Logo</div>
      <div class="gap-4 hidden lg:block">
        <router-link to="/" class="hover:underline"> Home </router-link>
        <a
          href="#explore"
          active-class="underline"
          class="hover:underline">
          Explore
        </a>
        <a href="#about" class="hover:underline"> About </a>
        <a href="#contact" class="hover:underline"> Contact </a>
      </div>

      <div class="flex gap-4">
        <Button class="rounded-full" variant="outline">
          <RouterLink to="/login">Sign In</RouterLink>
        </Button>
        <Button class="rounded-full" variant="default">
          <RouterLink to="/register">Sign Up</RouterLink>
        </Button>
      </div>
    </nav>
  </header>
  <div
    class="lg:hidden flex items-center justify-between gap-4 w-full bg-white py-4 px-4 fixed bottom-0 rounded-l-4xl rounded-r-4xl">
    <router-link
      to="/"
      class="flex flex-col items-center justify-center hover:underline text-center">
      <House />
      <span class="text-xs text-muted-foreground mt-1">Home</span>
    </router-link>
    <a
      href="#explore"
      class="flex flex-col items-center justify-center hover:underline text-center">
      <Telescope />
      <span class="text-xs text-muted-foreground mt-1">Explore</span>
    </a>
    <a
      href="#about"
      class="flex flex-col items-center justify-center hover:underline">
      <Info />
      <span class="text-xs text-muted-foreground mt-1">About</span>
    </a>
    <a
      href="#contact"
      class="flex flex-col items-center justify-center hover:underline">
      <Send />
      <span class="text-xs text-muted-foreground mt-1">Contact</span>
</a>
    <a
      href="#profile"
      class="flex flex-col items-center justify-center hover:underline">
      <CircleUser />
      <span class="text-xs text-muted-foreground mt-1">Profile</span>
    </a>
  </div>
</template>
<script setup lang="ts">
import { RouterLink } from "vue-router";
import { Button } from "@/components/ui/button";
import { House, Telescope, Info, Send, CircleUser } from "lucide-vue-next";


</script>
