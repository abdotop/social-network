import { createResolver } from '@nuxt/kit'
const { resolve } = createResolver(import.meta.url)

const ONE_DAY = 60 * 60 * 24 * 1000;
const ONE_WEEK = ONE_DAY * 7;

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  ssr: false,
  app: {
    head: {
      link: [
        { rel: "stylesheet", href: "https://cdn.jsdelivr.net/npm/boxicons@latest/css/boxicons.min.css" },
        { rel: 'stylesheet', href: 'https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/css/select2.min.css' },
      ],
      script: [
        { src: 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js' },
        { src: 'https://cdnjs.cloudflare.com/ajax/libs/uikit/3.15.14/js/uikit.min.js' },
        { src: "https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/js/select2.min.js" },
        { type: "module", src: "https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.esm.js" },
        { src: "https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.js" }
      ],
    }
  },
  css: [
    '~/assets/css/style.css',
  ],
  modules: [
    '@nuxt/ui',
    '@vueuse/nuxt',
    // '@unocss/nuxt',
    '@nuxtjs/color-mode',
    '@pinia/nuxt',
    '@nuxt/image',
    'nuxt-icon',
    '@samk-dev/nuxt-uikit3'
  ],
  colorMode: {
    classSuffix: '',
  },
  typescript: {
    includeWorkspace: true,
  },
  imports: {
    dirs: [
      resolve('./composables'), '~/composables',
      resolve('./stores'), '~/stores',
      resolve('~/types'), '~/types',
      resolve('~/api'), '~/api',
      resolve('~/utils'), '~/utils'
    ],
  },
  image: {
    domains: ['localhost:8081'],
  },
  pinia: {
    storesDirs: ['~/stores/**', '#/stores/**', '@/stores/**'],
  },
  runtimeConfig: {
    domain: process.env.BACKEND_URL || 'http://localhost:8081',
    cookieName: process.env.COOKIE_NAME || "__social_cookie",
    cookieSecret: process.env.COOKIE_SECRET || "secret",
    cookieExpires: parseInt(process.env.COOKIE_REMEMBER_ME_EXPIRES || ONE_DAY.toString(), 10), // 1 day
    cookieRememberMeExpires: parseInt(process.env.COOKIE_REMEMBER_ME_EXPIRES || ONE_WEEK.toString(), 10), // 7 days
  },
  nitro: {
    experimental: {
      websocket: true,
      tasks: true,
    },
  },
})
