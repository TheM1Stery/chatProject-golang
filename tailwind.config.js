/** @type {import('tailwindcss').Config} */
export default {
  content: ["./routes/**/*.templ"],
  theme: {
    extend: {
      gridTemplateRows: {
        layout: "auto 50px",
      },
    },
  },
  plugins: [require("daisyui")],
};
