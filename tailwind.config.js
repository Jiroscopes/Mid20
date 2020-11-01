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
            backgroundImage:{
                'd2c':"url('/d2c.jpg')",
            },
            inset: {
                circleLines: '1400px',
                circle2XL: '1780px',
                circle2LG: '2900px',
                circle2MD: '2800px',
                lg: '130px',
                xl: '430px',
            },
            height: {
                portImg: '300px',
                circleMD: '70%',
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
        width: ['hover'],
        backgroundImage: ({ before }) => before(['active']),
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
