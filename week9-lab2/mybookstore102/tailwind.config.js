/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  safelist: [
    'bg-sky-100',
    'bg-rose-100', 
    'bg-lime-100',
    'bg-indigo-100'
  ],
  theme: {
    extend: {
      colors: {
        'bookstore-primary': '#2d5a4d',
        'bookstore-secondary': '#5fe9bc',
        viridian: {
          50: '#f0f9f4',
          100: '#dcf2e4', 
          200: '#bce4cc',
          300: '#8dcfa8',
          400: '#57b37e',
          500: '#33996a',
          600: '#267d54',
          700: '#1f6544',
          800: '#1b5037',
          900: '#17422f',
        }
      },
      fontFamily: {
        'sans': ['Prompt', 'sans-serif'],
      }
    },
  },
  plugins: [],
}