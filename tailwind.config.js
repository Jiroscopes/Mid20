module.exports = {
    prefix: '',
    important: false,
    separator: ':',
    theme: {
        fontFamily: {
            'Nunito': ['nunito', 'sans-serif'],
            'OpenSans': ['"Open Sans"', 'sans-serif']
        },
        extend: {
            inset: {
                lg: '130px',
                xl: '430px',
            },
            height: {
                wwd: '425px',
                wwdHover: '450px',
            },
            width: {
                wwdHover: '450px',
            },
            padding: {
                xxl: '13rem',
            },
            colors: {
                orange: '#EE6F57',
                blue: '#1F3C88',
                darkBlue: '#070D59',
                white: '#F6F5F5'
            },
        }
   },
    variants: {
        height: ['hover'],
        width: ['hover']
    },
   plugins: [],
   purge: {
    enabled: process.env.NODE_ENV === 'production',
    content: [
        'components/**/*.vue',
        'layouts/**/*.vue',
        'pages/**/*.vue',
        'plugins/**/*.js',
        'nuxt.config.js',
        // TypeScript
        'plugins/**/*.ts',
        'nuxt.config.ts'
    ]
  }
 }
