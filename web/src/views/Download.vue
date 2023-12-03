<template>
    <main-nav :title="title" />
    <div class="download">
        <h1>{{ content.title }}</h1>
        <h2>{{ content.subtitle }}</h2>
        <div class="content-block" v-for="(paragraph, index) in content.paragraphs" :key="index">
            <p>{{ paragraph }}</p>
        </div>
        <div class="content-block" v-for="(link, index) in links" :key="index">
            <h1 style="color: #007BFF;">{{ link.title }}</h1>
            <h1><a :href="link.url" target="_blank">{{ link.name }}</a></h1>
            <h2><a :href="link.url" target="_blank">{{ link.otherName }}</a></h2>
        </div>
    </div>
</template>

<script>
export default {
    name: "Download",
    data() {
        return {
            content: {
                //title: 'AiMo3D插件下载',
                //subtitle: '欢迎来到AiMo社区。',
                paragraphs: [
                    
                ],
            },
            links: [
            {
                    //title: '插件下载链接：',
                    name: 'AiMo3D Rhino插件安装包',
                    url: 'https://release.zhai-zhai.com/s/k5h6',
                    otherName: '',
                    otherUrl: '',
                },
                {
                    //title: '插件教程文档',
                    name: '插件安装和使用教程',
                    url: 'https://tutorial.zhai-zhai.com',
                    otherName: '',
                    otherUrl: '',
                },
                {
                    title: '插件配套课程：',
                    name: '《教AI帮你做设计，居然一不小心学会了代码》',
                    url: 'https://rhinostudio.cn/my/course/3075',
                    otherName: '',
                    otherUrl: '',
                },
            ]
        }
    }
};
</script>

<style scoped>
.download {
    text-align: center;
    padding: 20px;
    background-color: #f9f9f9;
    border-radius: 10px;
    box-shadow: 0 0 10px rgba(0,0,0,0.1);
}

.download h1, .download h2 {
    color: #333;
    font-weight: bold;
}

.download .content-block {
    background-color: #fff;
    padding: 10px;
    margin: 10px 0;
    border-radius: 5px;
    box-shadow: 0 0 5px rgba(0,0,0,0.1);
}

.download p {
    color: #666;
}

.download a {
    color: #007BFF;
    text-decoration: none;
}

.download a:hover {
    color: #0056b3;
}
</style>