function toggleFlip(card) {
    if (card.style.transform === 'rotateX(180deg)') {
        card.style.transform = 'rotateX(0deg)';
    } else {
        card.style.transform = 'rotateX(180deg)';
    }
}