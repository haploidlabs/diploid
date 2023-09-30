// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  devtools: { enabled: true },
  modules: ["@pinia/nuxt", "@nuxtjs/google-fonts", "nuxt-lucide-icons"],
  imports: {
    dirs: ["store"],
  },
  css: ['~/assets/css/tailwind.css'],
  app: {
    head: {
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' }
      ],
      noscript: [
        { children: 'JavaScript is required' }
      ]
    }
  },
  pinia: {
    autoImports: [
      'defineStore',
      ['defineStore', 'definePiniaStore'],
    ],
  },
  googleFonts: {
    families: {
      Inter: [400, 500, 600, 700],
    }
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
});
