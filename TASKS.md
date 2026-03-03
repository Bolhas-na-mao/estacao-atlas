# Tarefas — Fase 1: O Despertar na Biblioteca

**Labels:** `sprite` `código` `som`

---

## HERO-01 · Criar sprites do herói
**Labels:** `sprite`

O herói precisa de sprites desenhados manualmente no Aseprite para substituir os assets de IA atuais. Devem cobrir idle e caminhada nas quatro direções (norte, sul, leste, oeste).

**Definição de Pronto:**
- Sprite sheet única com idle e caminhada para as 4 direções
- Exportado em `.png` compatível com o sistema atual
- Visualmente consistente (proporção, paleta de cores)

---

## HERO-02 · Refatorar lógica de desenho para sprite sheet única
**Labels:** `código`

Atualmente o herói usa arquivos separados para idle e caminhada (`hero_idle`, `hero_moving`), o que é ineficiente. A lógica de animação deve ser centralizada em uma única sprite sheet.

**Definição de Pronto:**
- Um único arquivo de sprite sheet carregado em memória
- `Draw` seleciona o frame correto com base no estado (idle/walking) e direção
- Comportamento visual idêntico ao anterior
- Arquivos antigos removidos

---

## MAP-01 · Criar sprite do piso da biblioteca
**Labels:** `sprite`

Tile de piso para as salas internas da biblioteca. Deve remeter a um ambiente antigo, empoeirado e literário.

**Definição de Pronto:**
- Tile desenhado no Aseprite, tamanho adequado para tiling
- Exportado em `.png`
- Funciona visualmente quando repetido em grid

---

## MAP-02 · Criar as duas primeiras salas da biblioteca (somente piso)
**Labels:** `código`

Criar as duas primeiras salas da biblioteca usando apenas o tile de piso, sem objetos. Objetivo é validar o sistema de câmera, colisões de borda e transição entre salas antes de adicionar assets visuais.

**Definição de Pronto:**
- Duas salas renderizadas com tile de piso
- Herói se move dentro dos limites das salas (sem sair pelos cantos)
- Transição entre as duas salas funciona corretamente

---

## MAP-03 · Criar lógica de mapa (limites e transição entre salas)
**Labels:** `código`

Implementar o sistema que define os limites de cada sala e gerencia a transição do herói ao atingir uma borda/porta.

**Definição de Pronto:**
- Herói não atravessa paredes ou limites do mapa
- Ao atingir a área de transição, a sala correta é carregada
- Sistema funciona para as salas existentes e é extensível para novas salas

---

## MAP-04 · Criar assets da biblioteca (objetos e decoração)
**Labels:** `sprite`

Desenhar os elementos visuais internos da biblioteca: mesas, janelas, livros, estantes, paredes, divisória entre parede e parte escura (interior sem luz direta).

**Definição de Pronto:**
- Assets desenhados no Aseprite com paleta e estilo consistentes entre si
- Exportados em `.png`
- Inclui: mesa, estante, livro(s), janela, parede, divisória parede/escuridão

---

## MAP-05 · Adicionar assets à biblioteca
**Labels:** `código`

Posicionar os assets criados em MAP-04 nas salas da biblioteca já construídas em MAP-02.

**Dependência:** MAP-04 concluída.

**Definição de Pronto:**
- Assets visualmente posicionados e sem sobreposições incorretas
- Colisão básica com objetos sólidos (herói não atravessa estantes/mesas)
- Salas com a ambientação esperada para a Fase 1

---

## MAP-06 · Criar terceira sala da biblioteca (porta de saída e porta do porão)
**Labels:** `código` `sprite`

A terceira sala conecta o interior da biblioteca à saída para o mundo exterior e ao acesso ao porão. É aqui que o Guardião das Sílabas bloqueia a entrada do porão. Inclui as duas portas visíveis.

**Definição de Pronto:**
- Sala renderizada com piso e assets existentes
- Duas portas visíveis: saída (mundo exterior) e entrada do porão
- Transições para as salas adjacentes funcionando
- Posição do Golem definida em frente à porta do porão

---

## SFX-01 · Adicionar som de passos
**Labels:** `som`

Reproduzir um efeito sonoro de passos enquanto o herói estiver em movimento.

**Definição de Pronto:**
- Som toca somente durante movimento
- Não toca durante idle
- Volume adequado (não intrusivo)

---

## NPC-01 · Criar sprite do Guardião das Sílabas (Golem)
**Labels:** `sprite`

O Guardião é um golem quadrado feito de livros empilhados. Para a Fase 1, apenas a direção sul (de frente para o jogador) é necessária.

**Definição de Pronto:**
- Sprite desenhado no Aseprite, direção sul
- Exportado em `.png`
- Visualmente remete a livros/blocos empilhados formando uma figura

---

## NPC-02 · Criar retrato do Guardião (estilo 3x4)
**Labels:** `sprite`

Ilustração do rosto/busto do Guardião no estilo retrato 3x4, exibida na caixa de diálogo durante a conversa.

**Definição de Pronto:**
- Imagem desenhada no Aseprite no formato/proporção definidos para retratos no jogo
- Exportada em `.png`
- Reconhecível como o Guardião

---

## DIALOG-01 · Criar caixa de diálogo (texto + retrato)
**Labels:** `sprite` `código`

Caixa de diálogo visual exibida durante a conversa com um NPC. Deve ter espaço para o texto da fala e o retrato (3x4) de quem está falando.

**Definição de Pronto:**
- Caixa renderizada na tela com área para texto e área para retrato
- Layout definido e reutilizável para outros NPCs
- Exibida corretamente sobre o cenário

---

## DIALOG-02 · Criar lógica de diálogo (iniciar, avançar, sair)
**Labels:** `código`

Implementar o fluxo de diálogo: aproximar-se do NPC → pressionar espaço para iniciar → espaço para avançar linhas → última linha fecha o diálogo.

**Definição de Pronto:**
- Diálogo inicia ao pressionar espaço perto do NPC
- Cada pressionamento avança uma linha
- Ao fim das linhas, o diálogo fecha automaticamente
- Movimento do herói bloqueado durante o diálogo

---

## DIALOG-03 · Criar animação de texto aparecendo (typewriter)
**Labels:** `código`

O texto da fala deve aparecer caractere por caractere, simulando uma máquina de escrever.

**Definição de Pronto:**
- Texto aparece progressivamente a uma velocidade agradável
- Pressionar espaço durante a animação exibe o texto completo instantaneamente
- Um único pressionamento de espaço não pode ao mesmo tempo completar o texto e avançar o diálogo: quando o espaço completa o texto, esse pressionamento é consumido e o avanço de linha só ocorre em um novo pressionamento (tecla solta e pressionada novamente)
- Funciona para qualquer string de diálogo

---

## DIALOG-04 · Adicionar som para o texto aparecendo
**Labels:** `som`

Efeito sonoro sincronizado com o aparecimento dos caracteres do diálogo.

**Dependência:** DIALOG-03 concluída.

**Definição de Pronto:**
- Som toca a cada caractere exibido (ou em intervalos regulares)
- Não toca quando o texto já está completo
- Volume adequado

---

## DIALOG-05 · Criar caixa de diálogo estilo pergunta/resposta
**Labels:** `código`

Variante da caixa de diálogo para os desafios do Guardião: exibe a pergunta e opções de resposta clicáveis/selecionáveis.

**Dependência:** DIALOG-01 concluída.

**Definição de Pronto:**
- Caixa exibe pergunta e lista de opções
- Jogador pode selecionar uma opção (teclado)
- Resposta selecionada é retornada para a lógica do jogo

---

## GAME-01 · Criar lógica da barra de Credibilidade
**Labels:** `código`

Implementar a mecânica central da Fase 1: barra de Credibilidade começa em 5, +1 para acerto, -1 para erro, vitória em 10, derrota em 0.

**Dependência:** DIALOG-05 concluída.

**Definição de Pronto:**
- Barra de Credibilidade renderizada na tela durante o desafio
- Valor atualiza corretamente ao acertar ou errar
- O valor de Credibilidade é sempre limitado ao intervalo 0..10: após qualquer modificação aplicar `max(0, min(10, novoValor))`
- Atingir 10 dispara o evento de vitória (verificar com o valor já limitado para evitar duplo disparo)
- Atingir 0 dispara o game over (verificar com o valor já limitado para evitar duplo disparo)

---

## GAME-02 · Criar tela de game over
**Labels:** `código` `sprite`

Tela exibida quando a Credibilidade chega a 0. Deve oferecer opção de tentar novamente (reinicia a fase) e voltar ao menu principal.

**Definição de Pronto:**
- Tela de game over exibida ao perder
- Botão/opção "Tentar novamente" reinicia a fase do zero
- Botão/opção "Menu" volta à tela inicial do launcher
- Visualmente adequada ao estilo do jogo

---

## GAME-03 · Criar animação da porta do porão abrindo e Golem saindo
**Labels:** `código` `sprite`

Animação de vitória: após atingir Credibilidade 10, a porta do porão se abre e o Golem sai do caminho, liberando a descida.

**Dependência:** GAME-01 concluída.

**Definição de Pronto:**
- Animação da porta do porão abrindo reproduzida após vitória
- Golem se move para fora do caminho ou desaparece com transição
- Após a animação, o herói pode entrar no porão

---

## MAP-07 · Criar sprite do porão
**Labels:** `sprite`

O porão é o destino final da Fase 1, acessado após vencer o Guardião. É onde a Pluma Lendária está guardada. Deve transmitir um ambiente escuro, empoeirado e claustrofóbico, distinto das salas superiores.

**Definição de Pronto:**
- Tile(s) e/ou cenário do porão desenhados no Aseprite
- Exportado em `.png`
- Visualmente distinto das salas superiores da biblioteca

---

## ITEM-01 · Criar sprite da Pluma Lendária
**Labels:** `sprite`

A Pluma Lendária é a recompensa da Fase 1, encontrada no porão após vencer o Guardião. Deve ser visualmente memorável e reconhecível como um item especial.

**Definição de Pronto:**
- Sprite desenhado no Aseprite
- Exportado em `.png`
- Visível no chão ou exibido como item coletável

---

## ITEM-02 · Criar lógica de inventário e coleta de item
**Labels:** `código`

Implementar o sistema básico de inventário e a interação de coleta: herói chega perto da Pluma → pressiona espaço → item é coletado e adicionado ao inventário.

**Dependência:** ITEM-01 concluída.

**Definição de Pronto:**
- Estrutura de inventário criada (mínimo: lista de itens coletados)
- Pluma some do mapa ao ser coletada
- Item registrado no inventário do herói
- Coleta só ocorre quando o herói está na área de interação

---

## ITEM-03 · Criar animação de coleta da Pluma
**Labels:** `código` `sprite`

Animação visual exibida no momento da coleta: a Pluma sobe brevemente e desaparece, dando feedback claro ao jogador.

**Dependência:** ITEM-02 concluída.

**Definição de Pronto:**
- Animação reproduzida ao coletar o item
- Pluma some do mapa após a animação terminar
- Animação não bloqueia o movimento do herói

---

## ITEM-04 · Adicionar som de coleta da Pluma
**Labels:** `som`

Efeito sonoro reproduzido no momento em que a Pluma é coletada.

**Dependência:** ITEM-02 concluída.

**Definição de Pronto:**
- Som toca uma vez no momento da coleta
- Volume adequado
- Transmite sensação de conquista/recompensa

---

*Total: 24 tarefas | `sprite`: 12 · `código`: 14 · `som`: 3*
