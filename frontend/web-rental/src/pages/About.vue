<template>
    <div class="grid grid-cols-1 lg:grid-cols-2 py-10 ">
        <div class="w-full max-w-[600px]">
           <h1 class="text-3xl font-bold">About me</h1>
           <h4 class="text-xl text-muted-foreground mt-2">Lorem ipsum dolor sit amet consectetur adipisicing elit.</h4>
           <p class="text-muted-foreground mt-4 leading-relaxed">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Dolores laborum blanditiis ad ipsa? Inventore in blanditiis optio aut quod, distinctio dolorem pariatur quia similique perspiciatis iste ipsum officiis laborum doloribus.
           </p>
           <Button variant="outline" class="w-full mt-2"> Learn more </Button>
        </div>
        <div class="w-full">
            <img src="https://placehold.co/1920x1080" alt="" class="w-full h-full object-cover">
        </div>
    </div>
</template>

<script setup lang="ts">
import{ Button } from '@/components/ui/button'
</script>