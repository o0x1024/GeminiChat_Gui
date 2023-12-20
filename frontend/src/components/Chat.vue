

<template>
  <div class="chat-box">
    <div class="gpt-title">
      <span class="gpt-title">Gemini Assist</span>
    </div>
    <div id="comment" class="comment">
      <p>Based on Gemini Pro API.</p>
    </div>

    <div id="result" class="result" ref="scrollableDiv">
      <div v-for="message in messages.list" :key="message.id" class="message">
        <template v-if="message.sender == 'human'">

          <strong><span style="color: rgb(48, 143, 10);">{{ message.sender }}</span></strong>
          : <span style="color: black;">

            <div style="background-color: rgb(235, 240, 240); line-height: 0px;border-radius: 10px; ">
              <p style="padding: 30px;">{{ message.text }}</p>
            </div>

            <!-- <div v-html="message.text"></div> -->
          </span>
          <p></p>
        </template>
        <template v-else>
          <strong><span style="color: rgb(223, 38, 38);">{{ message.sender }}</span></strong>
          : <span style="color: black; ">
            <p></p>
            <MdPreview :editorId="id" :modelValue="message.text"  style="background-color: rgb(235, 240, 240); line-height: 50px; padding: 0 10px;border-radius: 10px;"/>
            <MdCatalog :editorId="id" :scrollElement="scrollElement" />

          </span>
          <p></p>
        </template>


      </div>
    </div>

    <div id="input" class="input-box ">
      <input id="name" v-model="question" placeholder="say something" autocomplete="off" class="input" type="text"
        @keydown.enter="handleEnterKey" />
      <button class="btn" @click="send">send</button>
      <button class="clr" title="Clear" gen-slate-btn="" @click="clear"><svg xmlns="http://www.w3.org/2000/svg"
          width="1em" height="1em" viewBox="0 0 24 24">
          <path fill="currentColor"
            d="M8 20v-5h2v5h9v-7H5v7h3zm-4-9h16V8h-6V4h-4v4H4v3zM3 21v-8H2V7a1 1 0 0 1 1-1h5V3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v3h5a1 1 0 0 1 1 1v6h-1v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z">
          </path>
        </svg></button>
    </div>
  </div>
  <div class="setting">
    <button class="btn" @click="hidden_setting"> <svg focusable="false" class="" data-icon="setting" width="1em"
        height="1em" fill="currentColor" aria-hidden="true" viewBox="64 64 896 896">
        <path
          d="M924.8 625.7l-65.5-56c3.1-19 4.7-38.4 4.7-57.8s-1.6-38.8-4.7-57.8l65.5-56a32.03 32.03 0 009.3-35.2l-.9-2.6a443.74 443.74 0 00-79.7-137.9l-1.8-2.1a32.12 32.12 0 00-35.1-9.5l-81.3 28.9c-30-24.6-63.5-44-99.7-57.6l-15.7-85a32.05 32.05 0 00-25.8-25.7l-2.7-.5c-52.1-9.4-106.9-9.4-159 0l-2.7.5a32.05 32.05 0 00-25.8 25.7l-15.8 85.4a351.86 351.86 0 00-99 57.4l-81.9-29.1a32 32 0 00-35.1 9.5l-1.8 2.1a446.02 446.02 0 00-79.7 137.9l-.9 2.6c-4.5 12.5-.8 26.5 9.3 35.2l66.3 56.6c-3.1 18.8-4.6 38-4.6 57.1 0 19.2 1.5 38.4 4.6 57.1L99 625.5a32.03 32.03 0 00-9.3 35.2l.9 2.6c18.1 50.4 44.9 96.9 79.7 137.9l1.8 2.1a32.12 32.12 0 0035.1 9.5l81.9-29.1c29.8 24.5 63.1 43.9 99 57.4l15.8 85.4a32.05 32.05 0 0025.8 25.7l2.7.5a449.4 449.4 0 00159 0l2.7-.5a32.05 32.05 0 0025.8-25.7l15.7-85a350 350 0 0099.7-57.6l81.3 28.9a32 32 0 0035.1-9.5l1.8-2.1c34.8-41.1 61.6-87.5 79.7-137.9l.9-2.6c4.5-12.3.8-26.3-9.3-35zM788.3 465.9c2.5 15.1 3.8 30.6 3.8 46.1s-1.3 31-3.8 46.1l-6.6 40.1 74.7 63.9a370.03 370.03 0 01-42.6 73.6L721 702.8l-31.4 25.8c-23.9 19.6-50.5 35-79.3 45.8l-38.1 14.3-17.9 97a377.5 377.5 0 01-85 0l-17.9-97.2-37.8-14.5c-28.5-10.8-55-26.2-78.7-45.7l-31.4-25.9-93.4 33.2c-17-22.9-31.2-47.6-42.6-73.6l75.5-64.5-6.5-40c-2.4-14.9-3.7-30.3-3.7-45.5 0-15.3 1.2-30.6 3.7-45.5l6.5-40-75.5-64.5c11.3-26.1 25.6-50.7 42.6-73.6l93.4 33.2 31.4-25.9c23.7-19.5 50.2-34.9 78.7-45.7l37.9-14.3 17.9-97.2c28.1-3.2 56.8-3.2 85 0l17.9 97 38.1 14.3c28.7 10.8 55.4 26.2 79.3 45.8l31.4 25.8 92.8-32.9c17 22.9 31.2 47.6 42.6 73.6L781.8 426l6.5 39.9zM512 326c-97.2 0-176 78.8-176 176s78.8 176 176 176 176-78.8 176-176-78.8-176-176-176zm79.2 255.2A111.6 111.6 0 01512 614c-29.9 0-58-11.7-79.2-32.8A111.6 111.6 0 01400 502c0-29.9 11.7-58 32.8-79.2C454 401.6 482.1 390 512 390c29.9 0 58 11.6 79.2 32.8A111.6 111.6 0 01624 502c0 29.9-11.7 58-32.8 79.2z">
        </path>
      </svg></button>
    <template v-if="settingvisible">
      <input id="apikey" v-model="api_key" placeholder="apikey" autocomplete="off" class="input1" type="text" />
      <input id="proxy" v-model="proxy" placeholder="proxy" autocomplete="off" class="input2" type="text" />
    </template>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, onMounted, watch, nextTick } from 'vue'
import { SendMessage, GetConf, SyncConf } from '../../wailsjs/go/main/App'

import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';

const scrollElement = document.documentElement;
const id = 'preview-only';

interface Message {
  id: number
  sender: string
  text: string
}

let question = ref('')
let api_key = ref('AIzaSyAFeOAtvOra_i8BwvTeYw50WoIKU0y8wZY')
let proxy = ref('http://127.0.0.1:7890')
let settingvisible = ref(false)
let messages: { list: Message[] } = reactive({ list: [] });
const scrollableDiv = ref();

onMounted(() => {
  let mg: Message = { id: 1, sender: "robot", text: "How may I assist you?" }
  messages.list.push(mg)

  GetConf().then((res: any) => {
    api_key.value = res.APIKEY
    proxy.value = res.Proxy
  })
})

watch(messages.list, () => {
  scrollToBottom();
});


watch(api_key, () => {
  SyncConf(api_key.value, proxy.value)
});

watch(proxy, () => {
  SyncConf(api_key.value, proxy.value)

});


const scrollToBottom = () => {
  nextTick(() => {
    // 滚动到底部
    scrollableDiv.value.scrollTop = scrollableDiv.value.scrollHeight;
  });
  // scrollableDiv.value.scrollTop = scrollableDiv.value.scrollHeight;
};

function clear() {
  messages.list = []
  let mg: Message = { id: 1, sender: "robot", text: "How may I assist you?" }
  messages.list.push(mg)
}


const handleEnterKey = () => {
  send()
}


function send() {

  let mg: Message = { id: messages.list[messages.list.length - 1].id + 1, sender: "human", text: question.value }
  messages.list.push(mg)
  SendMessage(question.value, api_key.value, proxy.value).then((res: any) => {


    if (res.Code == 200) {
      let mg: Message = { id: messages.list[messages.list.length - 1].id + 1, sender: "robot", text: res.Message }
      messages.list.push(mg)
      console.log(mg)
      console.log(res.Message)
    } else {
      alert("报错了")
    }
  })
  question.value = ""
}


function hidden_setting() {
  settingvisible.value = !settingvisible.value
}


</script>


<style scoped>
.input-box {
  margin-top: 20px;
}


.chat-box {
  padding: 10px 20%;
}

.result {
  width: 800px;
  /* white-space: pre-line; */
  text-align: left;
  color: #babbbd;
  overflow-y: scroll;
  height: 640px;
  line-height: 20px;
  margin: 1.2rem auto;
}

.comment {
  text-align: left;
  color: #babbbd;
  height: 20px;
  line-height: 20px;
  margin: 1.2rem auto;
}


.gpt-title {
  margin-top: 10px;
  text-align: left;
  color: black;
  margin-right: 0.25rem;
  font-size: 1.8rem;
  line-height: 2rem;
  font-weight: 800;
}

.setting {
  position: absolute;
  left: 5px;
  bottom: 10px;
  /* width: 40px; */

}


.setting .btn {
  float: left;
  width: 40px;
  height: 30px;
  line-height: 35px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;

}

.setting .input1 {
  border: none;
  background-color: rgba(240, 240, 240, 1);
  height: 30px;
  width: 300px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  padding: 0 8px;
  cursor: pointer;
  margin-left: 10px;
}

.setting .input2 {
  margin-left: 10px;
  border: none;
  background-color: rgba(240, 240, 240, 1);
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  padding: 0 8px;
  cursor: pointer;
}


.input-box .btn {
  float: left;
  width: 90px;
  height: 45px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .clr {
  float: left;
  width: 40px;
  height: 45px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}


.input-box .btn:hover {
  float: left;
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  font-size: 16px;
  color: #656566;
  float: left;
  width: 60%;
  border: none;
  border-radius: 3px;
  outline: none;
  height: 45px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
