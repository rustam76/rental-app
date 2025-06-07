<script setup lang="ts">
import { Card, CardTitle } from '@/components/ui/card';

import { Button } from '@/components/ui/button';
import { ChevronLeft, ChevronRight } from 'lucide-vue-next'

const data = [
    {
        title: 'mobila pajero',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam, quae.',
        image: 'https://placehold.co/600x400',
        price: "Rp. 1.000.000"
    },
    {
        title: 'lorem',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam, quae.',
        image: 'https://placehold.co/600x400',
        price: "Rp. 2.000.000"
    },
    {
        title: 'ayam',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam, quae.',
        image: 'https://placehold.co/600x400',
        price: "Rp. 3.000.000"
    },
    {
        title: 'buku',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam, quae.',
        image: 'https://placehold.co/600x400',
        price: "Rp. 4.000.000"
    },
]

const buttonCheron = (value: string) => {
    if(value == 'left') {
        data.reverse()
    }
    if(value == 'right') {
        data.reverse()
    }
}

</script>

<template>
            <div class="flex items-center justify-between py-3">
            <h1>Discover Explore In Rental App</h1>
            <div class="flex items-center gap-4">
                <Button variant="outline" type="button" v-on:click="buttonCheron('left')">
                    <ChevronLeft class="w-4 h-4 mr-2"/>
                </Button>
                <Button variant="outline" type="button" v-on:click="buttonCheron('right')">
                    <ChevronRight class="w-4 h-4 mr-2"/>
                </Button>
            </div>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4" >
            <Card v-for="item in data" :key="item.title" >
                <CardContent>
                    <img :src="item.image" alt="" class="w-full h-full object-cover rounded-r-xl rounded-l-xl">
                </CardContent>
                <CardFooter class="px-4 py-4">
                    <CardTitle>{{ item.title }}</CardTitle>
                    <CardDescription>{{ item.description }}</CardDescription>
                    <CardDescription>{{ item.price }}</CardDescription>
                    <CardAction>
                        <Button variant="outline" class="w-full mt-2"> Learn more </Button>
                    </CardAction>
                </CardFooter>
            </Card>
        </div>
</template>