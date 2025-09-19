<script setup>

import { onMounted, onBeforeUnmount, ref } from 'vue'
import { Editor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Button from '../ui/Button.vue'
import Input from '../ui/Input.vue'

const editor = ref(null)

onMounted(() => {
  editor.value = new Editor({
    content: '',
    extensions: [
      StarterKit,
    ],
    editorProps: {
      attributes: {
        class: 'prose prose-sm sm:prose lg:prose-lg xl:prose-2xl mx-auto focus:outline-none p-4',
      },
    },
  })
})

onBeforeUnmount(() => {
  if (editor.value) {
    editor.value.destroy()
  }
})

</script>

<template>
  <main class="flex-1">
    <article class="editor max-w-[840px] h-full w-full flex flex-col justify-center mx-auto py-8">
      <Input name="article_title" type="text" placeholder="Title.."
        class="text-3xl border-none outline-none block pb-4" />
      <div class="editor-container">
        <editor-content :editor="editor" />
        <Button type="submit">Text</Button>
      </div>
    </article>
  </main>
</template>

<style>
.ProseMirror {
  cursor: text;
  min-height: 120px;
  --line-height: 24px;
  height: 100%;
  padding: 1rem;
  background-color: #2c2c2e;
  border: 1px solid #ffffff0f;
  border-radius: 24px;
  font-size: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, .1), 0 30px 100px rgba(192, 211, 255, .08);
}

.ProseMirror:focus {
  outline: none;
}

.ProseMirror p {
  margin: 0 0 1rem 0;
}

.ProseMirror h1,
.ProseMirror h2,
.ProseMirror h3,
.ProseMirror h4,
.ProseMirror h5,
.ProseMirror h6 {
  margin: 1.5rem 0 1rem 0;
  font-weight: bold;
}

.ProseMirror ul,
.ProseMirror ol {
  padding-left: 1.5rem;
  margin: 0 0 1rem 0;
}

.editor-container {
  position: relative;
}
</style>