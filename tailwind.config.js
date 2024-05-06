/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["**/*_templ.go", "**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
}

