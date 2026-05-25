import React, { KeyboardEvent, useState } from 'react';
import './App.css';

function TagInput() {
  const [inputValue, setInputValue] = useState<string>('');
  const [tags, setTags] = useState<string[]>([]);
  const [errorMessage, setErrorMessage] = useState<string>('');

  const addTag = () => {
    const trimmedValue = inputValue.trim();

    if (!trimmedValue) {
      setErrorMessage('Tag cannot be empty.');
      return;
    }

    if (tags.includes(trimmedValue)) {
      setErrorMessage('Duplicate tags are not allowed.');
      return;
    }

    setTags((currentTags) => [...currentTags, trimmedValue]);
    setInputValue('');
    setErrorMessage('');
  };

  const handleKeyDown = (event: KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      event.preventDefault();
      addTag();
    }
  };

  const handleRemoveTag = (tagToRemove: string) => {
    setTags((currentTags) => currentTags.filter((tag) => tag !== tagToRemove));
    setErrorMessage('');
  };

  return (
    <section className="tag-input-card">
      <div className="tag-input-header">
        <p className="tag-input-eyebrow">ReactJS & TypeScript</p>
        <h1>Simple Tag Input</h1>
        <p className="tag-input-description">
          Type a tag and press Enter to add it. Empty values and duplicates are rejected.
        </p>
      </div>

      <div className="tag-list" aria-live="polite">
        {tags.length === 0 ? (
          <p className="tag-list-empty">No tags added yet.</p>
        ) : (
          tags.map((tag) => (
            <span className="tag-pill" key={tag}>
              <span>{tag}</span>
              <button
                type="button"
                className="tag-remove-button"
                onClick={() => handleRemoveTag(tag)}
                aria-label={`Remove ${tag}`}
              >
                ×
              </button>
            </span>
          ))
        )}
      </div>

      <label className="tag-input-label" htmlFor="tag-input">
        Tag name
      </label>
      <input
        id="tag-input"
        className="tag-input-field"
        type="text"
        value={inputValue}
        onChange={(event) => {
          setInputValue(event.target.value);
          if (errorMessage) {
            setErrorMessage('');
          }
        }}
        onKeyDown={handleKeyDown}
        placeholder="Enter a tag and press Enter"
      />

      {errorMessage ? <p className="tag-input-error">{errorMessage}</p> : null}
    </section>
  );
}

export default TagInput;
