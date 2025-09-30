<script setup>
import { onMounted, onBeforeUnmount, ref } from 'vue'
import { Editor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import { useAlert } from '../../composables/useAlert'
import { useRoute, useRouter } from 'vue-router'
import Input from '../ui/Input.vue'
import Button from '../ui/Button.vue'

const editor = ref(null)
const title = ref('')

const route = useRoute()
const router = useRouter()
const currentEntryId = ref(null)

const { showAlert } = useAlert()

onMounted(async () => {
  editor.value = new Editor({
    content: '',
    extensions: [StarterKit],
    editorProps: {
      attributes: {
        class: 'prose prose-sm sm:prose lg:prose-lg xl:prose-2xl mx-auto focus:outline-none p-4',
      },
    },
  })

  if (route.path.startsWith('/edit/')) {
    const id = parseInt(route.params.id, 10)
    if (!isNaN(id)) {
      await loadEntry(id)
    }
  }
})



const addEntry = async () => {
  const titleValue = title.value.trim()
  const contentHTML = editor.value?.getHTML() || ''

  if (!titleValue || !editor.value?.state.doc.textContent.trim()) {
    showAlert("Поля не могут быть пустыми")
    return
  }

  // remove any time from here
  // create date with golang and time.Now on backend
  // avoit extra moves 

  try {
    await window.go.main.App.CreateEntry(titleValue, contentHTML)
    title.value = ""
    editor.value.commands.clearContent()
    showAlert("Запись добавлена")
  } catch (e) {
    console.error(e)
    showAlert("Ошибка при создании записи")
  }
}

const updateEntry = async () => {
  const titleValue = title.value.trim()
  const contentHTML = editor.value?.getHTML() || ''

  if (!titleValue || !editor.value?.state.doc.textContent.trim()) {
    showAlert("Поля не могут быть пустыми")
    return
  }

  try {
    await window.go.main.App.UpdateEntry({
      ID: currentEntryId.value,
      Title: titleValue,
      Text: contentHTML
    })
    showAlert("Запись обновлена")
    router.push({ name: 'article', params: { id: currentEntryId.value } })
  } catch (e) {
    console.error(e)
    showAlert("Ошибка при обновлении записи")
  }
}

const loadEntry = async (id) => {
  try {
    const entry = await window.go.main.App.GetEntry(id)
    title.value = entry.Title || ''
    editor.value.commands.setContent(entry.Text || '')
    currentEntryId.value = entry.ID
  } catch (e) {
    console.error('Ошибка загрузки записи:', e)
    showAlert('Не удалось загрузить запись')
    router.push({ name: 'home' })
  }
}
</script>

<template>
  <main class="flex-1">
    <article class="editor max-w-[840px] h-screen overflow-scroll w-full flex flex-col justify-center mx-auto py-8">
      <Input v-model="title" name="article_title" type="text" placeholder="Title.."
        class="text-3xl border-none outline-none block pb-4" />
      <div class="editor-container">
        <EditorContent :editor="editor" />
        <Button @click="route.path.startsWith('/edit/') ? updateEntry() : addEntry()" class="absolute right-6 bottom-6">
          {{ route.path.startsWith('/edit/') ? 'Сохранить' : 'Добавить' }}
        </Button>
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