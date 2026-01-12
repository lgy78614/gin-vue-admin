<!-- components/AceYamlEditor.vue -->
<template>
  <div class="ace-yaml-editor" ref="editorContainer" :style="{ height: height }"></div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import * as ace from 'ace-builds'
import 'ace-builds/src-noconflict/mode-yaml'
import 'ace-builds/src-noconflict/theme-monokai'
import 'ace-builds/src-noconflict/ext-language_tools'
import 'ace-builds/src-noconflict/ext-searchbox'
import 'ace-builds/src-noconflict/ext-beautify'
import 'ace-builds/src-noconflict/ext-error_marker'
import yaml from 'js-yaml'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  height: {
    type: String,
    default: '500px'
  },
  theme: {
    type: String,
    default: 'chrome'
  },
  readOnly: {
    type: Boolean,
    default: false
  },
  showLineNumbers: {
    type: Boolean,
    default: true
  },
  autoFormat: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue', 'change', 'init', 'error'])

const editorContainer = ref(null)
let editor = null
let errorMarker = null

// 设置 Ace 基础路径（关键步骤）
ace.config.set('basePath', 'https://cdn.jsdelivr.net/npm/ace-builds@1.27.0/src-noconflict/')
ace.config.setModuleUrl('ace/mode/yaml_worker', 'https://cdn.jsdelivr.net/npm/ace-builds@1.27.0/src-noconflict/worker-yaml.js')

const initEditor = () => {
  if (!editorContainer.value) return
  
  // 创建编辑器
  editor = ace.edit(editorContainer.value, {
    mode: 'ace/mode/yaml',
    theme: `ace/theme/${props.theme}`,
    value: props.modelValue,
    fontSize: 14,
    tabSize: 2,
    useSoftTabs: true,
    wrap: true,
    showLineNumbers: props.showLineNumbers,
    showGutter: true,
    showPrintMargin: false,
    highlightActiveLine: true,
    enableBasicAutocompletion: true,
    enableLiveAutocompletion: true,
    enableSnippets: true,
    readOnly: props.readOnly,
    maxLines: Infinity,
    minLines: 10,
    autoScrollEditorIntoView: true
  })

  // 配置语言工具
  editor.setOptions({
    enableBasicAutocompletion: true,
    enableLiveAutocompletion: true,
    enableSnippets: true
  })

  // 监听内容变化
  editor.on('change', () => {
    const value = editor.getValue()
    emit('update:modelValue', value)
    emit('change', value)
    
    // 实时验证 YAML 语法
    validateYaml(value)
  })

  // 添加命令
  editor.commands.addCommand({
    name: 'formatYaml',
    bindKey: { win: 'Ctrl-Shift-F', mac: 'Command-Shift-F' },
    exec: formatYaml
  })

  editor.commands.addCommand({
    name: 'validateYaml',
    bindKey: { win: 'Ctrl-Shift-V', mac: 'Command-Shift-V' },
    exec: validateYaml
  })

  // 初始化完成
  emit('init', editor)
  
  // 初始验证
  validateYaml(props.modelValue)
}

// 验证 YAML 语法
const validateYaml = (content = null) => {
  try {
    const yamlContent = content || editor.getValue()
    
    // 清除之前的错误标记
    clearErrorMarker()
    
    if (!yamlContent.trim()) return
    
    // 解析 YAML
    const parsed = yaml.load(yamlContent, {
      schema: yaml.DEFAULT_SAFE_SCHEMA,
      json: false
    })
    
    return { valid: true, data: parsed }
  } catch (error) {
    // 解析错误，标记错误位置
    markError(error)
    emit('error', error)
    return { valid: false, error }
  }
}

// 标记错误位置
const markError = (error) => {
  clearErrorMarker()
  
  if (!error.mark || !editor) return
  
  const line = error.mark.line + 1 // YAML 行号从0开始
  const column = error.mark.column
  
  // 创建会话
  const session = editor.getSession()
  
  // 添加错误标记
  errorMarker = session.addMarker(
    new ace.Range(line - 1, column, line - 1, column + 1),
    'ace_error-marker',
    'text'
  )
  
  // 添加注解
  session.setAnnotations([{
    row: line - 1,
    column: column,
    text: error.message,
    type: 'error'
  }])
  
  // 滚动到错误行
  editor.scrollToLine(line - 1, true, true, () => {})
}

// 清除错误标记
const clearErrorMarker = () => {
  if (errorMarker && editor) {
    editor.getSession().removeMarker(errorMarker)
    errorMarker = null
  }
  editor?.getSession().clearAnnotations()
}

// 格式化 YAML
const formatYaml = () => {
  try {
    const parsed = yaml.load(editor.getValue())
    const formatted = yaml.dump(parsed, {
      indent: 2,
      lineWidth: -1,
      noRefs: true,
      skipInvalid: true
    })
    editor.setValue(formatted, -1)
    clearErrorMarker()
  } catch (error) {
    markError(error)
  }
}

// 监听值变化
watch(() => props.modelValue, (newValue) => {
  if (editor && editor.getValue() !== newValue) {
    editor.setValue(newValue || '', -1)
  }
})

// 监听主题变化
watch(() => props.theme, (newTheme) => {
  if (editor) {
    editor.setTheme(`ace/theme/${newTheme}`)
  }
})

onMounted(() => {
  initEditor()
})

onBeforeUnmount(() => {
  if (editor) {
    editor.destroy()
    editor = null
  }
})

// 暴露方法给父组件
defineExpose({
  formatYaml,
  validateYaml,
  getEditor: () => editor,
  getValue: () => editor?.getValue(),
  setValue: (value) => editor?.setValue(value || '', -1)
})
</script>

<style scoped>
.ace-yaml-editor {
  width: 100%;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  overflow: hidden;
}

:deep(.ace_editor) {
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
}

:deep(.ace_error-marker) {
  position: absolute;
  background-color: rgba(255, 0, 0, 0.2);
  z-index: 5;
}

:deep(.ace_gutter-cell.ace_error) {
  background-image: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12"><circle cx="6" cy="6" r="5" fill="%23f56c6c"/></svg>');
  background-repeat: no-repeat;
  background-position: 2px center;
}
</style>