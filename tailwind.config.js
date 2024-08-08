/** @type {import('tailwindcss').Config} */
export default {
  content: ["./routes/**/*.templ"],
  theme: {
    extend: {
      gridTemplateRows: {
        layout: "1fr 50px",
      },
    },
  },
  plugins: [require("daisyui")],
};
