$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
  evento.preventDefault();
  console.log('Dentro da funcao criarUsuario');

  const senha = $('#senha').val();
  const confirmarSenha = $('#confirmar-senha').val();

  if (senha != confirmarSenha) {
    alert('As senhas não coincidem');
    return;
  }

  $.ajax({
    url: '/usuarios',
    method: 'POST',
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val(),
    },
  });
}
