import { createI18n } from "vue-i18n";
import en from "./en";
import zh from "./zh";

export default createI18n({
    locale:'zh',
    legacy:false,
    globalInjection:true,
    messages:{
        zh,
        en
    }
})