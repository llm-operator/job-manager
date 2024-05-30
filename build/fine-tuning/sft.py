# The original file is located at https://github.com/huggingface/trl/blob/main/examples/scripts/sft.py
#
# The code is slightly modified to remove the logic for using a rich progress bar.

import logging
import os
import argparse

import torch

from datasets import load_dataset
from transformers import AutoTokenizer, TrainingArguments, BitsAndBytesConfig
from peft import LoraConfig

from tqdm.rich import tqdm

from trl import (
    SFTTrainer,
    get_kbit_device_map,
)

# For progress bars.
tqdm.pandas()

if __name__ == "__main__":
    parser = argparse.ArgumentParser("sft.py", description="A script to train a model using SFT.")
    parser.add_argument("--model", help="Model path.", type=str)
    parser.add_argument("--dataset", help="Dataset path.", type=str)
    parser.add_argument("--output", help="Output path.", type=str)

    # TODO(kenji): Revisit the default values.
    parser.add_argument("--learning_rate", help="Learning rate.", default=2e-4, type=float, nargs="?")
    parser.add_argument("--num_train_epochs", help="Number of training epocs.", default=3, type=int, nargs="?")
    parser.add_argument("--per_device_train_batch_size", help="Batch size per training.", default=2, type=int, nargs="?")

    args = parser.parse_args()

    quantization_config = BitsAndBytesConfig(
        load_in_4bit=True,
        bnb_4bit_compute_dtype=torch.float16,
        bnb_4bit_quant_type="nf4"
    )

    model_kwargs = dict(
        # Which attention implementation to use; you can run --attn_implementation=flash_attention_2,
        # in which case you must install this manually by running `pip install flash-attn --no-build-isolation`
        attn_implementation=None,
        # Override the default `torch.dtype` and load the model under this dtype. If `auto` is passed,
        # the dtype will be automatically derived from the model's weights."
        torch_dtype='auto',
        # use_cache=False,
        device_map=get_kbit_device_map(),
        quantization_config=quantization_config,
    )

    # TODO(kenji): Revisit these parameters.
    training_args = TrainingArguments(
        output_dir=args.output,
        overwrite_output_dir=True,
        num_train_epochs=args.num_train_epochs,
        # batch size per device during training
        per_device_train_batch_size=args.per_device_train_batch_size,
        # number of steps before performing a backward/update pass
        gradient_accumulation_steps=2,
        # use gradient checkpointing to save memory
        gradient_checkpointing=True,
        # save checkpoint every epoch
        save_strategy="epoch",
        logging_steps=10,
        # learning rate, based on QLoRA paper
        learning_rate=args.learning_rate,
        # warmup ratio based on QLoRA paper
        warmup_ratio=0.03,
        # max gradient norm based on QLoRA paper
        max_grad_norm=0.3,
        # use constant learning rate scheduler
        lr_scheduler_type="constant",
        # use bfloat16 precision
        bf16=True,
        # use tf32 precision
        tf32=True,
        report_to="none",
    )

    raw_datasets = load_dataset(args.dataset)
    train_dataset = raw_datasets["train"]
    eval_dataset = raw_datasets["test"] if "test" in raw_datasets else None

    tokenizer = AutoTokenizer.from_pretrained("./base-model", use_fast=True)
    tokenizer.pad_token = tokenizer.eos_token

    # TODO(kenji): Revisit these parameters.
    peft_config = LoraConfig(
        r=16,
        lora_alpha=32,
        lora_dropout=0.05,
        bias="none",
        task_type="CAUSAL_LM",
        target_modules=None,
        modules_to_save=None,
    )

    trainer = SFTTrainer(
        model=args.model,
        model_init_kwargs=model_kwargs,
        args=training_args,
        train_dataset=train_dataset,
        eval_dataset=eval_dataset,
        # The name of the text field of the dataset, in case this is passed by a user,
        # the trainer will automatically create a `ConstantLengthDataset` based on the `dataset_text_field` argument.
        dataset_text_field=None,
        # Used only in case `dataset_text_field` is passed. This argument is used by the `ConstantLengthDataset`
        # to pack the sequences of the dataset.
        packing=False,
        # The maximum sequence length to use for the `ConstantLengthDataset`
        # and for automatically creating the Dataset.
        # Defaults to min of the smaller of the `tokenizer.model_max_length` and `1024`.
        max_seq_length=None,
        tokenizer=tokenizer,
        peft_config=peft_config,
        callbacks=None,
    )

    trainer.train()

    trainer.save_model(args.output)
