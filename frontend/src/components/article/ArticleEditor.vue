<script setup>
import { onMounted, onBeforeUnmount, ref } from 'vue'
import { Editor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'

import Input from '../ui/Input.vue'
import Alert from '../ui/Alert.vue'

const editor = ref(null)
const title = ref('')

// alert logic
const showAlert = ref(false)
const alertMessage = ref('')

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

const addEntry = async () => {
  const titleValue = title.value
  const contentHTML = editor.value.getHTML()


  if (titleValue.trim() === "" || !editor.value.state.doc.textContent.trim().length) {

    alertMessage.value = "Поля не могут быть пустыми"
    showAlert.value = true
    setTimeout(() => {
      showAlert.value = false
    }, 3000)
    return false

  }

  // remove any time from here
  // create date with golang and time.Now on backend
  // avoit extra moves 
  try {
    await window.go.main.App.CreateEntry(titleValue, contentHTML)
    title.value = ""
    editor.value.commands.clearContent()

    alertMessage.value = "Запись добавлена"
    showAlert.value = true
    setTimeout(() => {
      showAlert.value = false
    }, 3000)
  } catch (e) {
    console.log(e.message)
  }
}


</script>

<template>
  <main class="flex-1">
    <article class="editor max-w-[840px] h-full w-full flex flex-col justify-center mx-auto py-8">
      <Input v-model="title" name="article_title" type="text" placeholder="Title.."
        class="text-3xl border-none outline-none block pb-4" />
      <div class="editor-container">
        <editor-content :editor="editor" />
        <button @click.prevent="addEntry()" type="submit"
          class="hover:bg-neutral-700 absolute right-4 bottom-4 h-[34px] px-[12px] border-[1px] rounded-[18px] border-white/12 flex items-center justify-center cursor-pointer text-[13px] transition-colors duration-200">
          Button
        </button>
      </div>
    </article>
    <Transition name="alert">
      <Alert v-if="showAlert">{{ alertMessage }}</Alert>
    </Transition>
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

.alert-enter-active,
.alert-leave-active {
  transition: all 0.3s ease;
}

.alert-enter-from,
.alert-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}
</style>