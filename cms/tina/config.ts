import { defineConfig } from "tinacms";

// Your hosting provider likely exposes this as an environment variable
const branch = process.env.HEAD || process.env.VERCEL_GIT_COMMIT_REF || "main";

export default defineConfig({
  branch,
  clientId: process.env.CLIENT_ID, // Get this from tina.io
  token: process.env.TINA_TOKEN, // Get this from tina.io

  build: {
    outputFolder: "admin",
    publicFolder: "public",
  },
  media: {
    tina: {
      mediaRoot: "",
      publicFolder: "public",
    },
  },
  schema: {
    collections: [
      {
        name: "pages",
        label: "Pages",
        path: "content/pages",
        fields: [
          {
            type: "string",
            name: "heroTitle",
            label: "Hero Title",

            required: true,
          },

          {
            type: "string",
            name: "subtitle",
            label: "Hero Subtitle",
          },
          {
            type: "string",
            name: "aboutTitle",
            label: "About Title",
          },
          { type: "rich-text", name: "aboutBio", label: "About Bio" },
          { type: "string", name: "projectsTitle", label: "Projects Title" },
          {
            type: "object",
            name: "project",
            label: "Project",
            list: true,

            fields: [
              { type: "string", name: "title", label: "Title" },
              { type: "rich-text", name: "description", label: "Description" },
              { type: "string", name: "siteLink", label: "Site Link" },
              { type: "string", name: "githubLink", label: "Github Link" },
              { type: "image", name: "image", label: "Image" },
            ],
          },
        ],
        format: "md",
      },
    ],
  },
});
