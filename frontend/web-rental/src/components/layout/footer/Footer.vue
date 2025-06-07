<template>
     <footer class="bg-secondary text-center p-4 text-sm ">
      &copy; {{ date.getFullYear() }} Rental. All rights reserved
    </footer>
</template>

<script setup lang="ts">
    const date = new Date();

</script>