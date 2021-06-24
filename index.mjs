import fetch from 'node-fetch'
import fs from 'fs'
import { exec } from 'child_process'

// Main generator
function Generate(API_NAME = '') {

  return fetch(`https://eu.api.ovh.com/1.0/${API_NAME}.json`)
  .then(res => res.json())
  .then(api => {
    // Models and interfaces
    const models = Object.keys(api.models)

    Object
    .entries(api.models)
    .sort((a, b) => {
      return a[0].localeCompare(b[0])
    })
    .map(([name, model]) => {
      var path = name.split('.')
      path.splice(-1, 1) // drop struct name

      var apiTs = `package ${API_NAME}\n\n// GENERATED SDK for ${API_NAME} API\n\n`

      // Model
      if (model.properties) {
        apiTs += `// ${model.description || ''}\ntype ${model.id} struct {\n`

        Object.entries(model.properties).forEach(([propName, prop]) => {
          apiTs += property(propName, prop)
        })

        apiTs += '}\n'
      }

      // Enum
      if (model.enum) {
        switch (model.enumType) {
          case 'string':
            console.log(model)
            apiTs += `// ${model.description || ''}\ntype ${model.id} string\n\n`

            apiTs += 'var (\n'
            apiTs +=  model.enum.sort((a, b) => b.localeCompare(a)).map(item => {
              return `\t${enumName(model.id, item)} ${model.id} = "${item}"`
            }).join('\n')

            apiTs += '\n)'
            break
          case 'long':
            console.log(model)
            apiTs += `// ${model.description || ''}\ntype ${model.id} float64\n\n`

            apiTs += 'var (\n'
            apiTs +=  model.enum.sort((a, b) => b - a).map(item => {
              return `\t${enumName(model.id, item)} ${model.id} = ${item}`
            }).join('\n')

            apiTs += '\n)'

            break
          default:
            console.log('unknown enum type', model)
        }
      }

      //const fullpath = ['./'].concat(path).join('/').toLowerCase()
      const fullpath = `./${API_NAME}`
      mkdir(fullpath)
      fs.writeFileSync(fullpath+`/${model.id.toLowerCase()}.go`, apiTs, { flag: 'w'})
    })

    // Render this API
  })
  .catch(err => console.error('Err', err))
}

fetch('https://eu.api.ovh.com/1.0/')
.then(res => res.json())
.then(apis => {
  apis.apis
  .filter(({path}) => ['/cloud', '/auth', '/me'].includes(path))
  .map(({path}) => Generate(path.replace(/^\//g, '')))
})
.then(() => {
  exec('go fmt ./...', (error, stdout, stderr) => {
    if (error) {
      console.log(`error: ${error.message}`)
      return
    }
    if (stderr) {
      console.log(`stderr: ${stderr}`)
      return
    }
    //console.log(`stdout: ${stdout}`)
  })
})


function mkdir(path, root) {

  var dirs = path.split('/'), dir = dirs.shift(), root = (root || '') + dir + '/';

  try { fs.mkdirSync(root); }
  catch (e) {
    //dir wasn't made, something went wrong
    if(!fs.statSync(root).isDirectory()) throw new Error(e);
  }

  return !dirs.length || mkdir(dirs.join('/'), root);
}

String.prototype.firstUppercase = function() {
  return this.charAt(0).toUpperCase() + this.slice(1)
}

function property(propName = '', prop) {
  console.log(prop)
  var { canBeNull, required } = prop

  prop.type = mapType(prop.type)

  propName = propName.replace(/-/g, '')

  return `\t${propName.firstUppercase()}\t${canBeNull? '*' : ''}${prop.type}\t\`json:"${propName}${required? '': ',omitempty'}"\`\n`
}

function enumName(id = '', str = '') {
  if (id.endsWith('Enum')) {
    //id = id.replace(/Enum/g, '')
  }

  return id + str
    .toUpperCase()
    .replace(/-/g, '_')
    .replace(/:/g, '_')
    .replace(/\./g, '_')
    .replace(/\//g, '')
}

function  mapType(type = '') {
  if (type.endsWith('[]')) {
    return '[]' + mapType(type.replace('[]', ''))
  }

  switch (type) {
    case 'datetime':
    case 'date':
    case 'duration':
    case 'password':
    case 'ip':
    case 'ipv4':
    case 'ipv4Block':
    case 'ipv6Block':
    case 'ipv6':
    case 'ipBlock':
    case 'text':
    case 'uuid':
    case 'string':
      return 'string'
    case 'long':
      return 'int64'
    case 'double':
      return 'float64'
    case 'boolean':
      return 'bool'
    case 'T':
      return 'interface{}'
    default:
      if (type.includes('<')) {
        return 'interface{}'
      }

      if (type.includes('.')) { // sub struct
        let parts = type.split('.')

        return `${parts[parts.length-1]}`
      }

      console.warn('Unknown type', type)
      return 'interface{}'
  }
}
