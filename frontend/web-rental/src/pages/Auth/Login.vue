<script setup lang="ts">
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as zod from "zod";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import type { Login } from "@/types/auth";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

const formSchema = toTypedSchema(
  zod.object({
    email: zod.string().email(),
    password: zod.string().min(8),
  })
);

const form = useForm({
  validationSchema: formSchema,
});

const onsubmit = form.handleSubmit((e: Login) => {
  console.log(e);
});
</script>
<template>
  <div class="w-full lg:grid lg:min-h-screen lg:grid-cols-2">
    <div class="hidden bg-muted lg:block">
      <img
        src="https://placehold.co/1920x1080"
        alt="Image"
        width="1920"
        height="1080"
        class="h-full w-full object-cover dark:brightness-[0.2] dark:grayscale" />
    </div>
    <div class="flex items-center justify-center py-12">
      <div class="mx-auto grid w-[350px] gap-6">
        <div class="grid gap-2 text-center">
          <h1 class="text-3xl font-bold">Login</h1>
          <p class="text-balance text-muted-foreground">
            Enter your email below to login to your account
          </p>
        </div>
        <form @submit="onsubmit" class="grid gap-4">
          <FormField v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel for="email">Email</FormLabel>
              <FormControl>
                <Input
                  id="email"
                  placeholder="m@example.com"
                  v-bind="componentField"
                  required
                 />
              </FormControl>
              <FormMessage name="email"  />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="password">
            <FormItem>
              <FormLabel for="password">Password</FormLabel>
              <FormControl>
                <Input
                  id="password"
                  type="password"
                  required
                  v-bind="componentField" />
                  
              </FormControl>
              <FormMessage name="password" />
            </FormItem>
          </FormField>
          <Button type="submit" class="w-full"> Login </Button>
          <Button variant="outline" class="w-full"> Login with Google </Button>
        </form>
        <div class="mt-4 text-center text-sm">
          Don't have an account?
          <router-link to="/register" class="underline"> Sign Up </router-link>
        </div>
      </div>
    </div>
  </div>
</template>
