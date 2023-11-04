import { createApp } from "vue";
import "uno.css";
import App from "./App.vue";
import router from "./router";
import 'element-plus/dist/index.css';
import LoginCard from "@/components/LoginCard/LoginCard.vue"

// 引入elementPlus图标
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
	app.component(key, component);
}
// 引入pinia
import { createPinia } from "pinia";
const pinia = createPinia();
app.component("LoginCard", LoginCard);

app.use(pinia);
app.use(router);
app.mount("#app");
