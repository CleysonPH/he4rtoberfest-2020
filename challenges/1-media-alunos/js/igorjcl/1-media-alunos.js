const [nota1, nota2] = process.argv.slice(2);
const media = (Number(nota1) + Number(nota2)) / 2;

console.log(`Média -> ${media}`);
