import React from 'react';
import { fireEvent, render, screen } from '@testing-library/react';
import App from './App';

test('renders tag input heading', () => {
  render(<App />);

  expect(screen.getByText(/simple tag input/i)).toBeInTheDocument();
  expect(screen.getByLabelText(/tag name/i)).toBeInTheDocument();
});

test('adds and removes a tag', () => {
  render(<App />);

  const input = screen.getByLabelText(/tag name/i);

  fireEvent.change(input, { target: { value: 'golang' } });
  fireEvent.keyDown(input, { key: 'Enter', code: 'Enter', charCode: 13 });

  expect(screen.getByText('golang')).toBeInTheDocument();

  fireEvent.click(screen.getByRole('button', { name: /remove golang/i }));

  expect(screen.queryByText('golang')).not.toBeInTheDocument();
});

test('rejects empty and duplicate tags', () => {
  render(<App />);

  const input = screen.getByLabelText(/tag name/i);

  fireEvent.change(input, { target: { value: '   ' } });
  fireEvent.keyDown(input, { key: 'Enter', code: 'Enter', charCode: 13 });
  expect(screen.getByText(/tag cannot be empty/i)).toBeInTheDocument();

  fireEvent.change(input, { target: { value: 'react' } });
  fireEvent.keyDown(input, { key: 'Enter', code: 'Enter', charCode: 13 });
  expect(screen.getByText('react')).toBeInTheDocument();

  fireEvent.change(input, { target: { value: 'react' } });
  fireEvent.keyDown(input, { key: 'Enter', code: 'Enter', charCode: 13 });
  expect(screen.getByText(/duplicate tags are not allowed/i)).toBeInTheDocument();
});
