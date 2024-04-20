/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      backgroundImage: {
        'default': 'linear-gradient(to top, #93cf30, #e2ac20, #c62222)',
      }
    },
  },
  plugins: [],
}

