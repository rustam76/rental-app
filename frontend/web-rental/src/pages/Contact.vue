<template>
    <div class="border border-rouded-4xl w-full h-[200px] mx-auto flex flex-col gap-4 items-center justify-center bg-primary rounded-4xl">
        <div class="grid gap-2 text-center">
            <h1 class="text-3xl font-bold text-white">Contact</h1>
            <p class="text-balance text-muted-foreground">
                Enter your email below to login to your account
              </p>
        </div>
        <div class="flex w-full max-w-xl items-center gap-1.5">
    <Input id="email" type="email" placeholder="Email" class="bg-white"/>
    <Button type="submit" variant="outline">
      Subscribe
    </Button>
  </div>
    </div>
</template>

<script setup lang="ts">
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
</script>