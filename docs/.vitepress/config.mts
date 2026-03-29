import { defineConfig } from 'vitepress';

export default defineConfig({
  title: 'VeloeraGen2',
  description: 'Veloera Community Edition Documentation',
  lang: 'zh-CN',
  base: '/VeloeraGen2/',
  cleanUrls: true,
  themeConfig: {
    nav: [
      { text: '首页', link: '/' },
      { text: '文档', link: '/guide/' },
      { text: 'API', link: '/api/' },
      { text: '致谢', link: '/thanks' },
      { text: 'GitHub', link: 'https://github.com/moehans-official/VeloeraGen2' },
    ],
    sidebar: {
      '/guide/': [
        {
          text: '指南',
          items: [
            { text: '文档总览', link: '/guide/' },
            { text: '快速开始', link: '/guide/getting-started' },
            { text: '部署指南', link: '/guide/deployment' },
          ],
        },
      ],
      '/api/': [
        {
          text: 'API',
          items: [{ text: '接口概览', link: '/api/' }],
        },
      ],
    },
    socialLinks: [{ icon: 'github', link: 'https://github.com/moehans-official/VeloeraGen2' }],
    footer: {
      message: 'VeloeraGen2 Documentation',
      copyright: `Copyright ${new Date().getFullYear()} VeloeraGen2`,
    },
  },
});
